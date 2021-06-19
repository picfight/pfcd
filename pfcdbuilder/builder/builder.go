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
	"regexp"
	"strconv"
	"strings"
)

func SortPackages(input string) ([]string, deps.DepsGraph){

	gomodlist := ListFiles(input, nil, ALL_CHILDREN, ut.Ext("mod"))
	root := fileops.Parent(fileops.Parent(fileops.Parent(input)))
	inputs := Relatives(root, gomodlist)


	graph := deps.DepsGraph{map[string]*deps.GoModHandler{}}
	for k, _ := range inputs {
		gomod := ReadGoMod(inputs[k], k)
		//pin.S("gomod", gomod)
		graph.Vertices[gomod.Tag] = gomod
	}

	sorted := ut.SortGraph(graph)

	//sorted = Resort(sorted, graph)

	return sorted, graph
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

func findLineWith(lines []string, s string) int {
	for i, e := range lines {
		if strings.Contains(e, s) {
			return i
		}
	}
	return -1
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