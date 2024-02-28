package graph

import "fmt"

func (g *Graph) BFS(val int) {
	visited := make(map[int]bool)

	visited[val] = true

	arr := []int{val}

	for len(arr) > 0 {
		v := arr[0]
		fmt.Printf("%d ", v)
		arr = arr[1:]

		for _,e := range g.Vertex[v] {
			if !visited[e] {
				arr = append(arr, e)
				visited[e] = true
			}
		} 
	}
}


func (g *Graph) DFS(val int,visited map[int]bool) {
	visited[val] = true

	fmt.Printf("%d ",val)

	for _,v := range g.Vertex[val] {
		if !visited[v] {
			g.DFS(v,visited)
		}
	}
}


func (g *Graph) IsPathExists(a,b int) (bool,int) {
	visited := make(map[int]bool)
	count := 0
	arr := []int{a}

	for len(arr) > 0 {
		val := arr[0]
		arr = arr[1:]
		visited[val] = true
		count++
		for _,v := range g.Vertex[val] {
			if v == b {
				return true,count
			}
			if !visited[v] {
				arr = append(arr, v)
				visited[v] = true
			}
		}
	}
	return false,0
}

