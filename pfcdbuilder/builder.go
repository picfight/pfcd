package main

import (
	"github.com/jfixby/coinknife"
	"github.com/jfixby/pin"
	"github.com/jfixby/pin/commandline"
	"github.com/jfixby/pin/fileops"
	"path/filepath"
	"strings"
)

func main() {

	set := &coinknife.Settings{
		PathToInputRepo:        `D:\PICFIGHT\src\github.com\decred\dcrd`,
		PathToOutputRepo:       `D:\PICFIGHT\src\github.com\picfight\pfcd`,
		DoNotProcessAnyFiles:   false,
		DoNotProcessSubfolders: true,
		FileNameProcessor:      nameGenerator,
		IsFileProcessable:      processableFiles,
		FileContentProcessor:   fileGenerator,
		IgnoredFiles:           ignoredFiles(),
		InjectorsPath:          filepath.Join("", "code_injections"),
	}

	coinknife.Build(set)

}

func nameGenerator(data string) string {
	return data
}

func fileGenerator(data string) string {
	//data = coinknife.Replace(data, "github.com/decred/dcrd", "github.com/picfight/dcrd")
	//data = coinknife.Replace(data, "decred/dcrd", "picfight/dcrd")
	return data
}

// ignoredFiles
func ignoredFiles() map[string]bool {
	ignore := make(map[string]bool)
	ignore[".git"] = true
	ignore["vendor"] = true
	ignore["builder"] = true
	ignore["pfcdbuilder"] = true
	ignore["picfightcoin"] = true
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
	//------------------------------
	if strings.HasSuffix(file, ".mod") {
		return true
	}
	if strings.HasSuffix(file, "api.proto") {
		return true
	}
	if strings.HasSuffix(file, ".pot") {
		return true
	}
	if strings.HasSuffix(file, ".gyp") {
		return true
	}
	if strings.HasSuffix(file, ".cc") {
		return true
	}
	if strings.HasSuffix(file, ".h") {
		return true
	}
	if strings.HasSuffix(file, "notes.sample") {
		return true
	}
	if strings.HasSuffix(file, ".desktop") {
		return true
	}
	if strings.HasSuffix(file, ".log") {
		return true
	}
	if strings.HasSuffix(file, "dcrd.service") {
		return true
	}
	if strings.HasSuffix(file, ".conf") {
		return true
	}
	if strings.HasSuffix(file, ".json") {
		return true
	}
	if strings.HasSuffix(file, ".py") {
		return true
	}
	if strings.HasSuffix(file, ".go") {
		return true
	}
	if strings.HasSuffix(file, ".tmpl") {
		return true
	}
	if strings.HasSuffix(file, ".js") {
		return true
	}
	if strings.HasSuffix(file, ".sh") {
		return true
	}
	if strings.HasSuffix(file, ".css") {
		return true
	}
	if strings.HasSuffix(file, ".lock") {
		return true
	}
	if strings.HasSuffix(file, "LICENSE") {
		return true
	}
	if strings.HasSuffix(file, "CONTRIBUTORS") {
		return true
	}
	if strings.HasSuffix(file, "Dockerfile") {
		return true
	}
	if strings.HasSuffix(file, "Dockerfile.alpine") {
		return true
	}
	if strings.HasSuffix(file, "CHANGES") {
		return true
	}
	if strings.HasSuffix(file, ".iml") {
		return true
	}
	if strings.HasSuffix(file, ".yml") {
		return true
	}
	if strings.HasSuffix(file, ".toml") {
		return true
	}
	if strings.HasSuffix(file, ".md") {
		return true
	}
	if strings.HasSuffix(file, ".xml") {
		return true
	}
	if strings.HasSuffix(file, ".gitignore") {
		return true
	}
	if strings.HasSuffix(file, ".editorconfig") {
		return true
	}
	if strings.HasSuffix(file, ".eslintignore") {
		return true
	}
	if strings.HasSuffix(file, ".stylelintrc") {
		return true
	}
	if strings.HasSuffix(file, "config") {
		return true
	}
	if strings.HasSuffix(file, ".html") {
		return true
	}
	if strings.HasSuffix(file, ".po") {
		return true
	}
	if strings.HasSuffix(file, ".less") {
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
