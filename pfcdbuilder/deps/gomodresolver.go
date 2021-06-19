package deps

import (
	"fmt"
	"github.com/jfixby/pin"
	"github.com/jfixby/pin/lang"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

type GoModResolver struct {
}

type GoModPath struct {
}

func (r *GoModResolver) ResolveGoModPath(root string) *GoModPath {
	panic("")
}

func (r *GoModResolver) ReadGoMod(gomodfilepath *GoModPath) *GoModHandler {
	panic("")
}

func GetXML(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("GET error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Status error: %v", resp.StatusCode)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("Read body: %v", err)
	}

	return string(data), nil
}

func ReadGoMod(tag GitTag) *GoModHandler {
	result := &GoModHandler{}

	//url := "https://" + tag.Package + "/releases/tag/" + tag.ReleaseTag + "/go.mod"

	gitorg := tag.GitOrg()
	repo := tag.GitRepo()

	url := "https://raw.githubusercontent.com/" + gitorg + "/" + repo + "/release-v1.6.2/go.mod"

	iData, err := GetXML(url)
	lang.CheckErr(err)

	//iData := fileops.ReadFileToString(i)
	lines := strings.Split(iData, "\n")

	indexM := findLineWith(lines, "module")
	if indexM == -1 { // no dependencies
		lang.ReportErr("")
	}
	{
		sr := strings.Split(lines[indexM], "module ")
		pin.AssertTrue("", len(sr) == 2)
		result.Tag = sr[1]
	}

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
		depp := Dependency{
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
		depp := Dependency{
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
