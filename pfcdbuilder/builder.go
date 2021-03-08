package main

import (
	"github.com/jfixby/coinknife"
	"github.com/jfixby/pin"
	"github.com/jfixby/pin/commandline"
	"github.com/jfixby/pin/fileops"
	"github.com/jfixby/pin/lang"
	"github.com/picfight/pfcd/pfcdbuilder/ut"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	input := GoPath("decred/dcrd")
	output := GoPath("picfight/pfcd")
	pin.D("input", input)
	pin.D("output", output)
	fileops.EngageDeleteSafeLock(true)
	ClearProject(output, ignoredFiles())

	//gomodlist := ListFiles(input, ignoredFiles(), ALL_CHILDREN, ut.FoldersOnly)
	//pin.D("process", gomodlist)

	gomodlist := ListFiles(input, nil, ALL_CHILDREN, ut.Ext("mod"))
	inputs := Relatives(input, gomodlist)
	outputs := map[string]string{}
	for k, _ := range inputs {
		outputs[k] = output + k
	}

	for k, _ := range inputs {
		ConvertGoMod(inputs[k], outputs[k])
	}

	//pin.D("inputs ", inputs)
	//pin.D("outputs", outputs)
}

func ConvertGoMod(i string, o string) {
	iData := fileops.ReadFileToString(i)
	lines := strings.Split(iData, "\n")
	index0 := findLineWith(lines, "require(")
	sl := lines[index0]
	pin.D("sl", sl)
	oData := iData
	fileops.WriteStringToFile(o, oData)
}

func findLineWith(lines []string, s string) int {
	for i, e := range lines {
		if strings.Contains(e, s) {
			return i
		}
	}
	return -1
}

func Relatives(root string, subfiles []string) map[string]string {
	result := map[string]string{}
	for _, e := range subfiles {
		key := e[len(root):len(e)]
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
	filter ut.FileFilter) []string {
	if fileops.IsFile(target) {
		lang.ReportErr("This is not a folder: %v", target)
	}

	files, err := ioutil.ReadDir(target)
	lang.CheckErr(err)
	result := []string{}
	for _, f := range files {
		fileName := f.Name()
		filePath := filepath.Join(target, fileName)

		if IgnoredFiles[fileName] {
			continue
		}
		if fileops.IsFolder(filePath) && children != DIRECT_CHILDREN {
			children := ListFiles(filePath, IgnoredFiles, children, filter)
			result = append(result, children...)
			continue
		}

		if fileops.IsFile(filePath) {
			if filter(filePath) {
				result = append(result, filePath)
			}
			continue
		}
	}
	if filter(target) {
		result = append(result, target)
	}
	lang.CheckErr(err)
	return result
}

func GoPath(git string) string {
	return filepath.Join(os.Getenv("GOPATH"), "src", "github.com", git)
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
	//ignore[".github"] = true
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
