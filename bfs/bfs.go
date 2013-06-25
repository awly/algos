// Breadth-first search
// BFS implementation with a simple test graph
// not tested on bigger graphs
package main

import (
	"fmt"
)

func main() {
	adj := map[string][]string{
		"a": []string{"b", "c"},
		"b": []string{"d", "c"},
		"c": []string{"z", "x"},
		"d": []string{"a", "x"},
	}
	bfs("a", adj)
}

func bfs(start string, adj map[string][]string) {
	front := []string{start}
	parent := make(map[string]string)
	level := make(map[string]int)
	level[start] = 0
	i := 1

	for len(front) > 0 {
		next := make([]string, 0)
		fmt.Print(i, "\t")
		for _, f := range front {
			fmt.Print(f, " ")
			for _, n := range adj[f] {
				if _, ok := level[n]; ok {
					continue
				}
				level[n] = i
				parent[n] = f
				next = append(next, n)
			}
		}
		fmt.Println()
		i++
		front = next
	}
}
