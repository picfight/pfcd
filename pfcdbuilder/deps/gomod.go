package deps

import (
	"github.com/jfixby/pin"
	"github.com/picfight/pfcd/pfcdbuilder/ut"
	"strings"
)

type Dependency struct {
	Import  string
	Version string
}

type GoModHandler struct {
	Tag          string
	Dependencies []Dependency
}

type DepsGraph struct {
	Vertices map[string]*GoModHandler
}

type DepsGraphVertex struct {
	g *DepsGraph
	v *GoModHandler
}

func (v DepsGraphVertex) ListChildren() []ut.Vertex {
	result := []ut.Vertex{}

	deps := v.v.Dependencies
	PREF := "github.com/decred/dcrd/"
	for _, dp := range deps {
		im := dp.Import
		if !strings.HasPrefix(im, PREF) {
			continue
		}
		key := im[len(PREF)-1:] + "/go.mod"
		cv := v.g.Vertices[key]
		pin.D("         key", key)
		pin.D("v.g.Vertices", v.g.Vertices)
		pin.AssertNotNil(key, cv)
		c := &DepsGraphVertex{v.g, cv}
		result = append(result, c)
	}
	return result
}

func (d DepsGraph) ListVertices() []ut.Vertex {
	result := []ut.Vertex{}
	for _, v := range d.Vertices {
		result = append(result, &DepsGraphVertex{&d, v})
	}
	return result
}
