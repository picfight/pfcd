package ut

import "github.com/jfixby/pin"

type Graph interface {
	ListVertices() []string
	ListChildrenForVertex(vertexID string) []string
}

func SortGraph(g Graph) []string {
	result := []string{}
	remaning := map[string]bool{}
	for _, e := range g.ListVertices() {
		remaning[e] = true
	}

	visited := &vertexset{map[string]bool{}}

	for len(remaning) > 0 {
		var last string = ""
		for k, v := range remaning {
			if v {
				last = k
				pin.AssertNotNil("", last)
				break
			}
		}
		delete(remaning, last)
		result = append(result, dfs(g, last, visited)...)
	}

	return result
}

type vertexset struct {
	set map[string]bool
}

func dfs(g Graph, currentVertex string, visited *vertexset) []string {
	result := []string{}

	if visited.set[currentVertex] {
		return result
	}

	children := g.ListChildrenForVertex(currentVertex)
	for _, c := range children {
		result = append(result, dfs(g, c, visited)...)
	}

	result = append(result, currentVertex)
	visited.set[currentVertex] = true
	return result
}
