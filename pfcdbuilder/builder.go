package main

import (
	"encoding/json"
	"github.com/jfixby/coinknife"
	"github.com/jfixby/pin"
	"github.com/jfixby/pin/commandline"
	"github.com/jfixby/pin/fileops"
	"github.com/jfixby/pin/lang"
	"github.com/picfight/pfcd/pfcdbuilder/builder"
	"github.com/picfight/pfcd/pfcdbuilder/deps"
	"github.com/picfight/pfcd/pfcdbuilder/policy"
	"github.com/picfight/pfcd/pfcdbuilder/replacer"
	"github.com/picfight/pfcd/pfcdbuilder/ut"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const POLICY_FILE = "convert.plc"

func main() {
	ROOT_D := "github.com/decred/dcrd"
	ROOT_P := "github.com/picfight/pfcd"
	input := builder.GoPath(ROOT_D)
	output := builder.GoPath(ROOT_P)
	//policies := "policies/"

	//pin.D("input", input)
	//pin.D("output", output)
	fileops.EngageDeleteSafeLock(true)
	builder.ClearProject(output, ignoredFiles())

	sorted, graph := builder.SortPackages(input)

	processor := deps.NewProcessor()

	pin.D("sorted", sorted)

	for _, tag := range sorted {
		vx := graph.Vertices[tag]
		//pin.D(tag, vx.Dependencies)
		plc := PolicyFor(tag, vx)

		if plc == nil {
			pin.D("no policy for", vx.Tag)
			pin.D(tag, vx.Dependencies)
			lang.ReportErr("no policy for %v", vx.Tag)
		}

		if plc.IsSkipProcessing() {
			pin.D("skip", vx.Tag)
			processor.AddSkippedPackage(vx.Tag)
			continue
		}

		if plc.IsConvertFiles() {
			processor.AddRedirect(tag, coinknife.Replace(tag, ROOT_D, ROOT_P))
			ConvertPackage(input, output, tag, vx, plc, processor)
		}
		//plc := ReadPolicy(vx, policies)
		//
		//puotput := output
		//if plc.RedirectPackageTo != "" {
		//	processor.AddRedirect(vx.Tag, plc.RedirectPackageTo)
		//}
		//if !plc.IsSkipProcessing() {
		//	if plc.RedirectPackageTo != "" {
		//		puotput = GoPath(plc.RedirectPackageTo)
		//		ClearProject(puotput, ignoredFiles())
		//	}
		//	ConvertPackage(vx, input, puotput, plc, processor)
		//} else {
		//	pin.D("skip", vx.Tag)
		//}
		pin.D("---------------------------------------------------------------------")
	}
}

func ConvertGoFile(input string, output string, tag string, vx *deps.GoModHandler, plc *policy.PackagePolicy, proc *deps.Processor) {

	I := builder.GoPath(vx.Tag + "/go.mod")
	O := coinknife.Replace(I, input, output)
	//output + vx.Tag + "/go.mod"
	pin.D(I, O)
	iData := fileops.ReadFileToString(I)
	for pi, po := range proc.Redirects() {
		iData = coinknife.Replace(iData, pi, po)
		pin.D("   replace "+pi, po)
	}

	pin.D(I, iData)

	fileops.WriteStringToFile(O, iData)

}

func ClearDestination(input string, output string, vx *deps.GoModHandler) {
	I := builder.GoPath(vx.Tag)
	O := coinknife.Replace(I, input, output)
	//output + vx.Tag + "/go.mod"
	os.MkdirAll(O, os.ModePerm)
	builder.ClearProject(O, ignoredFiles())
}

func ConvertPackage(input string, output string, tag string, vx *deps.GoModHandler, plc *policy.PackagePolicy, proc *deps.Processor) {
	ClearDestination(input, output, vx)
	ConvertGoFile(input, output, tag, vx, plc, proc)
	if 1 == 1 {
		return
	}
	//pin.D("   tag", vx.Tag)
	//pin.D(" input", input+vx.Tag)
	//pin.D("output", output+vx.Tag
	//)

	{
		//pin.D(p, string(file))

		lang.AssertValue("package name", plc.PackageName, vx.Tag)

		I := builder.GoPath(vx.Tag)

		gofiles := builder.ListFiles(I, nil, builder.DIRECT_CHILDREN, ut.OR(ut.Ext("go"), ut.Name("README.md")))
		short2long := builder.ShortenFileNames(gofiles)

		pfiles := map[string]policy.FilePolicy{}

		for _, f := range plc.Files {
			if !f.IsValid() {
				lang.ReportErr("Invalid policy: %v", f)
			}
			pfiles[f.FileName] = f
			if short2long[f.FileName] != "" && (f.FileName != "go.mod") {
				lang.ReportErr("File not found : %v", f.FileName)
			}
		}

		for f, _ := range short2long {
			x := pfiles[f]
			if x.FileName == "" {
				//i := short2long[f]
				//iData := fileops.ReadFileToString(i)
				//pin.D(f, iData)
				//lang.ReportErr("No policy for file : %v", f)
			}
			//o := strings.ReplaceAll(i, input, output)

			//ConvertFile(i,o)
		}

		if plc.IsRedirectPackage() {
			s := "go.mod"
			{
				i := filepath.Join(I, s)
				//o := strings.ReplaceAll(i, input, output)
				o := filepath.Join(output, s)
				iData := fileops.ReadFileToString(i)
				iData = coinknife.Replace(iData, vx.Tag, plc.RedirectPackageTo)
				fileops.WriteStringToFile(o, iData)
			}
		}

		//
		{
			//for s, i := range short2long {
			//	//o := strings.ReplaceAll(i, input, output)
			//	o := filepath.Join(output, s)
			//	fp := pfiles[s]
			//	ConvertFile(i, o, &fp, plc, vx)
			//}
		}
	}
}

func PolicyFor(tag string, vx *deps.GoModHandler) *policy.PackagePolicy {
	if tag == "github.com/decred/dcrd/bech32" {
		return &policy.PackagePolicy{SkipProcessing: policy.YES}
	}
	if tag == "github.com/decred/dcrd/certgen" {
		return &policy.PackagePolicy{SkipProcessing: policy.YES}
	}
	if tag == "github.com/decred/dcrd/crypto/blake256" {
		return &policy.PackagePolicy{SkipProcessing: policy.YES}
	}
	if tag == "github.com/decred/dcrd/crypto/ripemd160" {
		return &policy.PackagePolicy{SkipProcessing: policy.YES}
	}
	if tag == "github.com/decred/dcrd/dcrec" {
		return &policy.PackagePolicy{SkipProcessing: policy.YES}
	}
	if tag == "github.com/decred/dcrd/dcrec/edwards" {
		return &policy.PackagePolicy{SkipProcessing: policy.YES}
	}
	if tag == "github.com/decred/dcrd/lru" {
		return &policy.PackagePolicy{SkipProcessing: policy.YES}
	}
	if tag == "github.com/decred/dcrd/chaincfg/chainhash" {
		return &policy.PackagePolicy{
			ConvertFiles: policy.YES,
			UseInjectors: policy.YES,
		}
	}
	if tag == "github.com/decred/dcrd/chaincfg/" {
		return &policy.PackagePolicy{
			ConvertFiles: policy.YES,
			UseInjectors: policy.YES,
		}
	}
	return nil
}

func ReadPolicy(vx *deps.GoModHandler, policies string) *policy.PackagePolicy {
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

	return plc
}

func ConvertFile(i string, o string, policy *policy.FilePolicy, P *policy.PackagePolicy, vx *deps.GoModHandler, proc *deps.Processor) {
	//if !policy.IsValid() {
	//	lang.ReportErr("Invalid policy: %v", policy)
	//}

	//pin.D("Convert: ", policy.FileName)
	pin.D("Convert: ", i)
	//pin.D(i, o)
	//si := filepath.Base(i)
	//so := filepath.Base(o)
	//lang.AssertValue("i", si, policy.FileName)
	//lang.AssertValue("o", so, policy.FileName)

	{
		iData := fileops.ReadFileToString(i)
		//pin.D(iData)
		fileops.WriteStringToFile(o, iData)
	}

	if policy.IsUseAutoReplacer() {
		iData := fileops.ReadFileToString(i)
		iData = replacer.ReplaceAll(iData)

		for k, v := range proc.Redirects() {
			iData = coinknife.Replace(iData, k, v)
		}

		fileops.WriteStringToFile(o, iData)
	}
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
