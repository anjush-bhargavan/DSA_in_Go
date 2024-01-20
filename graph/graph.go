package graph

import (
	"fmt"
)

type Graph struct {
	Vertex map[int][]int
}

func(g *Graph) Insert(vertex,edge int, isBidirectional bool) {
	g.Vertex[vertex] = append(g.Vertex[vertex], edge)
	if isBidirectional {
		g.Vertex[edge] = append(g.Vertex[edge], vertex)
	}
}

func(g *Graph) RemoveVertex(vertex int) {
	if g.Vertex == nil {
		return
	}
	for k,v := range g.Vertex {
		
		for i := 0 ; i<len(v) ; i++ {
			if v[i] == vertex {
				// fmt.Println(v[:i],v[i+1:])
				g.Vertex[k] = append(v[:i],v[i+1:]... )
				break
			}
		}
	}
	delete(g.Vertex,vertex)
}


func (g *Graph) RemoveEdge(v1, v2 int,isBidirectional bool) {
	arr1 := g.Vertex[v1]
	for i := 0 ; i < len(arr1) ; i ++ {
		if arr1[i] == v2 {
			g.Vertex[v1] = append(arr1[:i],arr1[i+1:]... )
			break
		}
	}
	if isBidirectional {
		arr1 := g.Vertex[v2]
		for i := 0 ; i < len(arr1) ; i ++ {
			if arr1[i] == v1 {
				g.Vertex[v2] = append(arr1[:i],arr1[i+1:]... )
				break
			}
		}
	}
}

func (g *Graph) Display() {
	for key,values := range g.Vertex {
		fmt.Println(key,values)
	}
}

func (g *Graph) BFS(value int) {
	visited := make(map[int]bool)
	visited[value] = true
	arr := []int{value}

	for len(arr) > 0 {
		fmt.Println(arr[0])
		val := arr[0]
		arr = arr[1:]

		for _,v := range g.Vertex[val] {
			if !visited[v] {
				arr = append(arr, v)
				visited[v] = true
			}
		}
	}
}


func (g *Graph) DFS(value int) {
	visited := make(map[int]bool)
	visited[value] = true
	stack := []int{value}
	g.dfsHelper(visited,stack)
}

func (g *Graph) dfsHelper(visited map[int]bool, stack []int) {
	if len(stack) == 0 {
		return 
	}
	val := stack[len(stack)-1]
	fmt.Println(val)
	stack = stack[:len(stack)-1]

	for _,v := range g.Vertex[val] {
		if !visited[v] {
			visited[v] = true
			stack = append(stack, v)
			g.dfsHelper(visited,stack)
		}
	}
}

func (g *Graph) IsPathExists(a,b int) bool {
	visited := make(map[int]bool)
	visited[a] = true
	stack := []int{a}

	for len(stack) > 0 {

		val := stack[0]
		stack = stack[1:]
		if val == b {
			return true
		}

		for _,v := range g.Vertex[val] {
			if !visited[v] {
				visited[v] = true
				stack = append(stack, v)
			}
		}
	}
	return false
}