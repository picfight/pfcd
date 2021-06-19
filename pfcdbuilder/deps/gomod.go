package deps

import (
	"github.com/jfixby/pin"
	"strings"
)

type Dependency struct {
	Import  string
	Fork    int
	Version string
}

type GoModHandler struct {
	Tag          string
	Dependencies []Dependency
}

type DepsGraph struct {
	Vertices map[string]*GoModHandler
}

type GitTag struct {
	Package    string
	ReleaseTag string
}

func (t *GitTag) GitOrg() string {
	//"github.com/decred/dcrd/"
	array := strings.Split(t.Package,"/")
	pin.AssertTrue("", len(array) == 3)
	return array[1]
}

func (t *GitTag) GitRepo() string {
	array := strings.Split(t.Package,"/")
	pin.AssertTrue("", len(array) == 3)
	return array[2]
}

func (v DepsGraph) ListChildrenForVertex(vertexID string) []string {
	result := []string{}

	//dps := v.g.Vertices[v.Tag]
	dps := v.Vertices[vertexID]
	deps := dps.Dependencies
	DCRD_PREF := "github.com/decred/dcrd/"
	for _, dp := range deps {
		im := dp.Import
		if strings.HasPrefix(im, DCRD_PREF) {
			//key := im[len(DCRD_PREF)-1:]
			key := im
			cv := v.Vertices[key]
			if cv == nil {
				pin.D("missing key", key+" : "+dp.Version)
				pin.D("v.g.Vertices", v.Vertices)
				pin.AssertNotNil(key, cv)
			}
			//pin.D("v.g.Vertices", v.g.Vertices)
			pin.AssertNotNil(key, cv)
			result = append(result, cv.Tag)
			continue
		}
	}
	return result
}

func (d DepsGraph) ListVertices() []string {
	result := []string{}
	for _, v := range d.Vertices {
		result = append(result, v.Tag)
	}
	return result
}
