package graph

import "fmt"

type Graph struct {
	Vertex map[int][]int
}

func NewGraph() *Graph {
	return &Graph{Vertex: make(map[int][]int)}
}

func (g *Graph) Insert(v1, v2 int) {
	g.Vertex[v1] = append(g.Vertex[v1], v2)
	g.Vertex[v2] = append(g.Vertex[v2], v1)
}

func (g *Graph) RemoveEdge(v1, v2 int) {
	arr := g.Vertex[v1]
	for i := 0; i < len(arr); i++ {
		if arr[i] == v2 {
			g.Vertex[v1] = append(arr[:i], arr[i+1:]...)
			break
		}
	}
	arr = g.Vertex[v2]
	for i := 0; i < len(arr); i++ {
		if arr[i] == v1 {
			g.Vertex[v2] = append(arr[:i], arr[i+1:]...)
			break
		}
	}

}

func (g *Graph) RemoveVertex(v int) {
	arr := g.Vertex[v]

	for i := 0 ; i < len(arr); i++ {
		edges := g.Vertex[arr[i]]

		for j := 0; j < len(edges); j++ {
			if edges[j] == v {
				g.Vertex[arr[i]] = append(edges[:j],edges[j+1:]... )
				break
			}
		}
		
	}
	delete(g.Vertex,v)
}

func (g *Graph) Display() {
	fmt.Println(g.Vertex)
}
