package ut

import "github.com/jfixby/pin"

type Graph interface {
	ListVertices() []Vertex
}
type Vertex interface {
	ListChildren() []Vertex
}

func SortGraph(g Graph) []Vertex {
	result := []Vertex{}
	remaning := map[Vertex]bool{}
	for _, e := range g.ListVertices() {
		remaning[e] = true
	}

	visited := map[Vertex]bool{}

	for len(remaning) > 0 {
		var last Vertex = nil
		for k, v := range remaning {
			if v {
				last = k
				pin.AssertNotNil("", last)
				break
			}
		}
		delete(remaning, last)
		dfs(g, last, visited, result)
	}

	return result
}

func dfs(g Graph, current Vertex, visited map[Vertex]bool, result []Vertex) {
	if visited[current] {
		return
	}

	children := current.ListChildren()
	for _, c := range children {
		dfs(g, c, visited, result)
	}

	result = append(result, current)
	visited[current] = true
}
