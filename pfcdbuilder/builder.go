package main

import (
	"github.com/jfixby/coinknife"
	"path/filepath"
)

func main() {

	set := &coinknife.Settings{
		PathToInputRepo:      `D:\PICFIGHT\src\github.com\decred\dcrd`,
		PathToOutputRepo:     `D:\PICFIGHT\src\github.com\picfight\dcrd`,
		DoNotProcessAnyFiles: false,
		FileNameProcessor:    PicfightCoinFileNameGenerator,
		IsFileProcessable:    ProcessableFiles,
		FileContentProcessor: PicfightCoinFileGenerator,
		IgnoredFiles:         IgnoredFiles(),
		InjectorsPath:        filepath.Join("", "code_injections", "d"),
	}

	coinknife.Build(set)

}
