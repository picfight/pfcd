package main

import (
	"github.com/jfixby/pin/fileops"
	"github.com/picfight/coin_knife/eproc"
	"path/filepath"
	"strings"
)

func FixSecp256k1Checksum(targetProject string) {
	invalidParent := filepath.Join(targetProject, "btcec")
	invalid := filepath.Join(invalidParent, "secp256k1.go")
	fileops.Delete(invalid)

	batName := "checksum_update.bat"
	batTemplate := filepath.Join("assets", batName)
	batData := fileops.ReadFileToString(batTemplate)
	batData = strings.Replace(batData, "#TARGET_FOLDER#", invalidParent, -1)
	batFile := filepath.Join(batName)
	fileops.WriteStringToFile(batFile, batData)

	ext := &eproc.ExternalProcess{
		CommandName: batFile,
	}
	ext.Launch(true)
	ext.Wait()
}
