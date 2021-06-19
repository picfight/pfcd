package ut

import (
	"github.com/jfixby/pin/fileops"
	"path/filepath"
	"strings"
)

type FileFilter func(input string) bool

var All = func(filePath string) bool { return true }

var FilesOnly = func(filePath string) bool {
	return fileops.IsFile(filePath)
}

var FoldersOnly = func(filePath string) bool {
	return fileops.IsFolder(filePath)
}

func AND(a FileFilter, b FileFilter) FileFilter {
	return func(filePath string) bool {
		return a(filePath) && b(filePath)
	}
}

func OR(a FileFilter, b FileFilter) FileFilter {
	return func(filePath string) bool {
		return a(filePath) || b(filePath)
	}
}

func Ext(ext string) FileFilter {
	return func(filePath string) bool {
		return strings.HasSuffix(filePath, "."+ext)
	}
}

func Name(ext string) FileFilter {
	return func(filePath string) bool {
		return filepath.Base(filePath) == ext
	}
}
