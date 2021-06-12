package main

import (
	"encoding/json"
	"github.com/jfixby/coinknife"
	"github.com/jfixby/pin"
	"github.com/jfixby/pin/commandline"
	"github.com/jfixby/pin/fileops"
	"github.com/jfixby/pin/lang"
	"github.com/picfight/pfcd/pfcdbuilder/deps"
	"github.com/picfight/pfcd/pfcdbuilder/policy"
	"github.com/picfight/pfcd/pfcdbuilder/replacer"
	"github.com/picfight/pfcd/pfcdbuilder/ut"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

const POLICY_FILE = "convert.plc"

func main() {
	input := GoPath("github.com/decred/dcrd")
	output := GoPath("github.com/picfight/pfcd")
	policies := "policies/"

	//pin.D("input", input)
	//pin.D("output", output)
	fileops.EngageDeleteSafeLock(true)
	ClearProject(output, ignoredFiles())

	gomodlist := ListFiles(input, nil, ALL_CHILDREN, ut.Ext("mod"))
	root := fileops.Parent(fileops.Parent(fileops.Parent(input)))
	inputs := Relatives(root, gomodlist)
	outputs := map[string]string{}
	for k, _ := range inputs {
		outputs[k] = output + k
	}

	graph := deps.DepsGraph{map[string]*deps.GoModHandler{}}
	for k, _ := range inputs {
		gomod := ReadGoMod(inputs[k], k)
		//pin.S("gomod", gomod)
		graph.Vertices[gomod.Tag] = gomod
	}

	sorted := ut.SortGraph(graph)

	pin.D("sorted", sorted)
	for _, tag := range sorted {
		vx := graph.Vertices[tag]
		pin.D(tag, vx.Dependencies)
		ConvertPackage(vx, input, output, policies)
		pin.D("---------------------------------------------------------------------")

	}
	//pin.D("outputs", outputs)
}

func ConvertPackage(vx *deps.GoModHandler, input string, output string, policies string) {
	//pin.D("   tag", vx.Tag)
	//pin.D(" input", input+vx.Tag)
	//pin.D("output", output+vx.Tag
	//)
	{
		I := GoPath(vx.Tag + "/go.mod")
		//O := output + vx.Tag + "/go.mod"
		//pin.D(I, O)
		iData := fileops.ReadFileToString(I)
		pin.D(I, iData)
	}
	{

		//O := output + vx.Tag
	}

	{
		P, err := filepath.Abs(policies + vx.Tag)
		lang.CheckErr(err)
		p := filepath.Join(P, POLICY_FILE)

		if !fileops.FileExists(p) {
			err := os.MkdirAll(P, 0x777)
			lang.CheckErr(err)

			lang.ReportErr("policy file not found: %v", p)
		}

		file, err := ioutil.ReadFile(p)
		lang.CheckErr(err)
		plc := &policy.PackagePolicy{}
		err = json.Unmarshal([]byte(file), &plc)
		lang.CheckErr(err)
		//pin.D(p, string(file))

		lang.AssertValue("package name", plc.PackageName, vx.Tag)

		I := GoPath(vx.Tag)

		gofiles := ListFiles(I, nil, DIRECT_CHILDREN, ut.FilesOnly)
		set, short2long := ShortenFileNames(gofiles)

		pfiles := map[string]policy.FilePolicy{}

		for _, f := range plc.Files {
			if !f.IsValid() {
				lang.ReportErr("Invalid policy: %v", f)
			}
			pfiles[f.FileName] = f
			if !set[f.FileName] {
				lang.ReportErr("File not found : %v", f.FileName)
			}
		}

		for f, _ := range set {
			x := pfiles[f]
			if x.FileName == "" {
				i := short2long[f]
				iData := fileops.ReadFileToString(i)
				pin.D(f, iData)
				lang.ReportErr("No policy for file : %v", f)
			}
			//o := strings.ReplaceAll(i, input, output)

			//ConvertFile(i,o)
		}

		for s, _ := range set {
			i := short2long[s]
			o := strings.ReplaceAll(i, input, output)
			//iData := fileops.ReadFileToString(i)
			fp := pfiles[s]
			ConvertFile(i, o, &fp)
		}
	}

}

func ShortenFileNames(input map[string]bool) (set map[string]bool, short2long map[string]string) {
	set = map[string]bool{}
	short2long = map[string]string{}
	for k, v := range input {
		p := fileops.Parent(k)
		s := k[len(p)+1:]
		set[s] = v
		short2long[s] = k
	}
	return
}

func ConvertFile(i string, o string, policy *policy.FilePolicy) {
	if !policy.IsValid() {
		lang.ReportErr("Invalid policy: %v", policy)
	}

	pin.D("Convert: ", policy.FileName)
	pin.D(i, o)
	si := filepath.Base(i)
	so := filepath.Base(o)
	lang.AssertValue("i", si, policy.FileName)
	lang.AssertValue("o", so, policy.FileName)

	if policy.IsCopyAsIs() {
		iData := fileops.ReadFileToString(i)
		//pin.D(iData)
		fileops.WriteStringToFile(o, iData)
	}

	if policy.IsUseAutoReplacer() {
		iData := fileops.ReadFileToString(i)
		iData = replacer.ReplaceAll(iData)
		fileops.WriteStringToFile(o, iData)
	}
}



func ReadGoMod(i string, tag string) *deps.GoModHandler {
	result := &deps.GoModHandler{}

	mm := strings.Index(tag, "/go.mod")
	//pin.D("tag", tag)
	if mm == 0 {
		//pin.D("tag", tag)
		result.Tag = "/"

	} else {
		result.Tag = tag[:mm]
	}
	//pin.D("result.Tag", result.Tag)

	iData := fileops.ReadFileToString(i)
	lines := strings.Split(iData, "\n")
	index0 := findLineWith(lines, "require")
	if index0 == -1 { // no dependencies
		return result
	}

	sr := strings.Split(iData, "require")
	pin.AssertTrue("", len(sr) == 2)

	brBegin := strings.Index(sr[1], "(")
	if brBegin == -1 {
		tokens := strings.Split(sr[1][1:], " ")
		dep := tokens[0]
		ver := tokens[1][:len(tokens[1])-1]
		depp := deps.Dependency{
			Import:  Dep(dep),
			Fork:    Fork(dep),
			Version: ver,
		}
		result.Dependencies = append(result.Dependencies, depp)
		return result
	}
	brEnd := strings.Index(sr[1], ")")
	list := sr[1][brBegin+1+1 : brEnd]
	lines = strings.Split(list, "\n")
	lines = lines[0 : len(lines)-1]
	for _, l := range lines {
		tokens := strings.Split(l, " ")
		dep := tokens[0][1:]
		ver := tokens[1][:len(tokens[1])]
		depp := deps.Dependency{
			Import:  Dep(dep),
			Fork:    Fork(dep),
			Version: ver,
		}
		result.Dependencies = append(result.Dependencies, depp)
	}
	return result
}

func Fork(dep string) int {
	rxp := "v[0-9][0-9]*"
	var validID = regexp.MustCompile(rxp)

	i := strings.LastIndex(dep, "/")
	//prefix := dep[:i]
	postfix := dep[i+1:]

	if validID.MatchString(postfix) {
		ForkString := postfix[1:]
		f, err := strconv.Atoi(ForkString)
		lang.CheckErr(err)
		//pin.D(dep, f)
		return f
	}
	return -1
}

func Dep(dep string) string {
	rxp := "v[0-9][0-9]*"
	var validID = regexp.MustCompile(rxp)

	i := strings.LastIndex(dep, "/")
	prefix := dep[:i]
	postfix := dep[i+1:]

	if validID.MatchString(postfix) {
		//pin.D(dep, prefix)
		return prefix
	}
	//pin.D(dep)
	return dep
}

func findLineWith(lines []string, s string) int {
	for i, e := range lines {
		if strings.Contains(e, s) {
			return i
		}
	}
	return -1
}

func Relatives(root string, subfiles map[string]bool) map[string]string {
	result := map[string]string{}
	for e, _ := range subfiles {
		key := e[len(root)+1 : len(e)]
		result[key] = e
	}
	return result
}

const ALL_CHILDREN = true
const DIRECT_CHILDREN = !ALL_CHILDREN

func ListFiles(
	target string,
	IgnoredFiles map[string]bool,
	children bool,
	filter ut.FileFilter) map[string]bool {
	if fileops.IsFile(target) {
		lang.ReportErr("This is not a folder: %v", target)
	}

	files, err := ioutil.ReadDir(target)
	lang.CheckErr(err)
	result := map[string]bool{}
	for _, f := range files {
		fileName := f.Name()
		filePath := filepath.Join(target, fileName)
		filePath = strings.ReplaceAll(filePath, "\\", "/")
		if IgnoredFiles[fileName] {
			continue
		}
		if fileops.IsFolder(filePath) && children != DIRECT_CHILDREN {
			children := ListFiles(filePath, IgnoredFiles, children, filter)
			//result = append(result, children...)
			result = putAll(result, children)
			continue
		}

		if fileops.IsFile(filePath) {
			if filter(filePath) {
				//result = append(result, filePath)
				result[filePath] = true
			}
			continue
		}
	}
	if filter(target) {
		//result = append(result, target)
		result[target] = true
	}
	lang.CheckErr(err)
	return result
}

func putAll(result map[string]bool, children map[string]bool) map[string]bool {
	for k, v := range children {
		result[k] = v
	}
	return result
}

func GoPath(git string) string {
	return strings.ReplaceAll(filepath.Join(os.Getenv("GOPATH"), "src", git), "\\", "/")
}

func ClearProject(target string, ignore map[string]bool) {
	pin.D("clear", target)
	files, err := ioutil.ReadDir(target)
	lang.CheckErr(err)

	for _, f := range files {
		fileName := f.Name()
		filePath := filepath.Join(target, fileName)
		if ignore[fileName] {
			pin.D("  skip", filePath)
			continue
		}
		pin.D("delete", filePath)
		err := os.RemoveAll(filePath)
		lang.CheckErr(err)
	}
	pin.D("")

}

func main_old() {
	root := `D:\PICFIGHT\src\github.com\decred\dcrd`
	ignored := ignoredFiles()
	subfolders := fileops.ListFiles(root, fileops.FoldersOnly, fileops.DIRECT_CHILDREN)
	fileops.EngageDeleteSafeLock(true)
	for _, input := range subfolders {

		p := fileops.PathToArray(input)
		fileName := p[len(p)-1]
		if ignored[fileName] {
			continue
		}

		output := filepath.Join(`D:\PICFIGHT\src\github.com\picfight\pfcd`, fileName)

		ignore := make(map[string]bool)
		set := &coinknife.Settings{
			PathToInputRepo:        input,
			PathToOutputRepo:       output,
			DoNotProcessAnyFiles:   false,
			DoNotProcessSubfolders: false,
			FileNameProcessor:      nameGenerator,
			IsFileProcessable:      processableFiles,
			FileContentProcessor:   fileGenerator,
			IgnoredFiles:           ignore,
			InjectorsPath:          filepath.Join("", "code_injections"),
		}

		coinknife.Build(set)
	}
}

func nameGenerator(data string) string {
	//data = coinknife.Replace(data, "decred/dcrd", "picfight/pfcd")
	return data
}

func fileGenerator(data string) string {
	data = coinknife.Replace(data, "decred/dcrd", "picfight/pfcd")
	//data = coinknife.Replace(data, "github.com/decred/dcrd", "github.com/picfight/dcrd")
	//data = coinknife.Replace(data, "decred/dcrd", "picfight/dcrd")
	return data
}

// ignoredFiles
func ignoredFiles() map[string]bool {
	ignore := make(map[string]bool)
	ignore[".git"] = true
	ignore[".gitignore"] = true
	ignore[".idea"] = true
	//ignore["rpctest"] = true
	//ignore["vendor"] = true
	//ignore["docs"] = true
	//ignore["cmd"] = true
	//ignore["builder"] = true
	ignore["pfcdbuilder"] = true
	//ignore["pfcregtest"] = true
	//ignore["picfightcoin"] = true

	ignore["go.mod"] = true
	ignore["go.sum"] = true
	return ignore
}

// processableFiles
func processableFiles(file string) bool {
	if strings.HasSuffix(file, ".png") {
		return false
	}
	if strings.HasSuffix(file, ".jpg") {
		return false
	}
	if strings.HasSuffix(file, ".jpeg") {
		return false
	}
	if strings.HasSuffix(file, ".exe") {
		return false
	}
	if strings.HasSuffix(file, ".svg") {
		return false
	}
	if strings.HasSuffix(file, ".ico") {
		return false
	}
	if strings.HasSuffix(file, ".bin") {
		return false
	}
	if strings.HasSuffix(file, ".bin") {
		return false
	}
	if strings.HasSuffix(file, ".db") {
		return false
	}
	if strings.HasSuffix(file, ".bz2") {
		return false
	}
	if strings.HasSuffix(file, ".gz") {
		return false
	}
	if strings.HasSuffix(file, ".hex") {
		return false
	}
	if strings.HasSuffix(file, ".mp4") {
		return false
	}
	if strings.HasSuffix(file, ".gif") {
		return false
	}
	if strings.HasSuffix(file, ".ttf") {
		return false
	}
	if strings.HasSuffix(file, ".icns") {
		return false
	}
	if strings.HasSuffix(file, ".woff") {
		return false
	}
	if strings.HasSuffix(file, ".woff2") {
		return false
	}
	if strings.HasSuffix(file, ".eot") {
		return false
	}
	if strings.HasSuffix(file, ".sum") {
		return false
	}
	//-
	if strings.HasSuffix(file, "api.proto") {
		return false
	}
	if strings.HasSuffix(file, ".pot") {
		return false
	}
	if strings.HasSuffix(file, ".gyp") {
		return false
	}
	if strings.HasSuffix(file, ".cc") {
		return false
	}
	if strings.HasSuffix(file, ".h") {
		return false
	}
	if strings.HasSuffix(file, "notes.sample") {
		return false
	}
	if strings.HasSuffix(file, ".desktop") {
		return false
	}
	if strings.HasSuffix(file, ".log") {
		return false
	}
	if strings.HasSuffix(file, "pfcd.service") {
		return false
	}
	if strings.HasSuffix(file, ".conf") {
		return false
	}
	if strings.HasSuffix(file, ".json") {
		return false
	}
	if strings.HasSuffix(file, ".py") {
		return false
	}
	if strings.HasSuffix(file, ".tmpl") {
		return false
	}
	if strings.HasSuffix(file, ".js") {
		return false
	}
	if strings.HasSuffix(file, ".sh") {
		return false
	}
	if strings.HasSuffix(file, ".css") {
		return false
	}
	if strings.HasSuffix(file, ".lock") {
		return false
	}
	if strings.HasSuffix(file, "LICENSE") {
		return false
	}
	if strings.HasSuffix(file, "CONTRIBUTORS") {
		return false
	}
	if strings.HasSuffix(file, "Dockerfile") {
		return false
	}
	if strings.HasSuffix(file, "Dockerfile.alpine") {
		return false
	}
	if strings.HasSuffix(file, "CHANGES") {
		return false
	}
	if strings.HasSuffix(file, ".iml") {
		return false
	}
	if strings.HasSuffix(file, ".yml") {
		return false
	}
	if strings.HasSuffix(file, ".toml") {
		return false
	}
	if strings.HasSuffix(file, ".md") {
		return false
	}
	if strings.HasSuffix(file, ".xml") {
		return false
	}
	if strings.HasSuffix(file, ".gitignore") {
		return false
	}
	if strings.HasSuffix(file, ".editorconfig") {
		return false
	}
	if strings.HasSuffix(file, ".eslintignore") {
		return false
	}
	if strings.HasSuffix(file, ".stylelintrc") {
		return false
	}
	if strings.HasSuffix(file, "config") {
		return false
	}
	if strings.HasSuffix(file, ".html") {
		return false
	}
	if strings.HasSuffix(file, ".po") {
		return false
	}
	if strings.HasSuffix(file, ".less") {
		return false
	}

	//------------------------------
	if strings.HasSuffix(file, ".mod") {
		return true
	}
	if strings.HasSuffix(file, ".go") {
		return true
	}

	pin.E("Unknown file type", file)
	return false
}

func fixSecp256k1Checksum(targetProject string) {
	invalidParent := filepath.Join(targetProject, "btcec")
	invalid := filepath.Join(invalidParent, "secp256k1.go")
	fileops.Delete(invalid)

	batName := "checksum_update.bat"
	batTemplate := filepath.Join("assets", batName)
	batData := fileops.ReadFileToString(batTemplate)
	batData = strings.Replace(batData, "#TARGET_FOLDER#", invalidParent, -1)
	batFile := filepath.Join(batName)
	fileops.WriteStringToFile(batFile, batData)

	ext := &commandline.ExternalProcess{
		CommandName: batFile,
	}
	ext.Launch(true)
	ext.Wait()
}
