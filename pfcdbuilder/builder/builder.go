package builder

import (
	"github.com/jfixby/pin"
	"github.com/jfixby/pin/fileops"
	"github.com/jfixby/pin/lang"
	"github.com/picfight/pfcd/pfcdbuilder/deps"
	"github.com/picfight/pfcd/pfcdbuilder/ut"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func SortPackages(gomodresolver *deps.GoModResolver, rootd string, gomodfilepath *deps.GoModPath) []deps.Dependency {
	result := &[]deps.Dependency{}
	deps := map[string]*[]deps.Dependency{}
	CollectPackage(gomodresolver, deps, rootd, gomodfilepath)

	for k, v := range deps {
		pin.D(k, v)
	}

	return *result
}

func CollectImport(gomodresolver *deps.GoModResolver, deps map[string]*[]deps.Dependency, rootd string, dep deps.Dependency) {
	gomodfilepath := gomodresolver.ResolveGoModPath(dep.Import)
	CollectPackage(gomodresolver, deps, rootd, gomodfilepath)
}

func CollectPackage(gomodresolver *deps.GoModResolver, deps map[string]*[]deps.Dependency, rootd string, gomodfilepath *deps.GoModPath) {
	gomod := gomodresolver.ReadGoMod(gomodfilepath)
	for _, dep := range gomod.Dependencies {
		if strings.HasPrefix(dep.Import, rootd) {
			//pin.D("", dep)
			CollectImport(gomodresolver, deps, rootd, dep)
		}
		//pin.D("", dep)

	}
	deps[gomod.Tag] = &gomod.Dependencies
}

func Swap(sorted []string, x int, y int) {
	sorted[x], sorted[y] = sorted[y], sorted[x]
}

func Relatives(root string, subfiles map[string]bool) map[string]string {
	result := map[string]string{}
	for e, _ := range subfiles {
		key := e[len(root)+1 : len(e)]
		result[key] = e
	}
	return result
}

func IsBigger(x string, y string, graph deps.DepsGraph) bool {
	if len(graph.ListChildrenForVertex(x)) == len(graph.ListChildrenForVertex(y)) {
		return x > y
	}
	return len(graph.ListChildrenForVertex(x)) > len(graph.ListChildrenForVertex(y))
}

func Resort(sorted []string, graph deps.DepsGraph) []string {

	N := len(sorted)
	swap := true
	for {
		for i := 0; i < N-1; i++ {
			if IsBigger(sorted[i], sorted[i+1], graph) {
				Swap(sorted, i, i+1)
				swap = true
			}
		}
		if !swap {
			break
		}
		swap = false
	}

	return sorted
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

func ShortenFileNames(input map[string]bool) (short2long map[string]string) {
	short2long = map[string]string{}
	for k, _ := range input {
		s := filepath.Base(k)
		short2long[s] = k
	}
	return
}
