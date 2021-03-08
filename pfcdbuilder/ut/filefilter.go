package ut

import (
	"github.com/jfixby/pin/fileops"
	"strings"
)

type FileFilter func(input string) bool

var AllFiles = func(filePath string) bool { return true }

var FoldersOnly = func(filePath string) bool {
	return fileops.IsFolder(filePath)
}

func AND(a FileFilter, b FileFilter) FileFilter {
	return func(filePath string) bool {
		return a(filePath) && b(filePath)
	}
}

func Ext(ext string) FileFilter {
	return func(filePath string) bool {
		return strings.HasSuffix(filePath, "."+ext)
	}
}
