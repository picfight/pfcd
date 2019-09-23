package main

import (
	"github.com/picfight/coin_knife"
	"path/filepath"
)

func main() {

	set := &coin_knife.Settings{
		PathToInputRepo:      `D:\PICFIGHT\src\github.com\btcsuite\btcd`,
		PathToOutputRepo:     `D:\PICFIGHT\src\github.com\btcziggurat\btcd`,
		DoNotProcessAnyFiles: false,
		FileNameProcessor:    PicfightCoinFileNameGenerator,
		IsFileProcessable:    ProcessableFiles,
		FileContentProcessor: PicfightCoinFileGenerator,
		IgnoredFiles:         IgnoredFiles(),
		InjectorsPath:        filepath.Join("", "code_injections", "d"),
	}

	coin_knife.Build(set)

}