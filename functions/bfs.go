package lemino

import "fmt"

// BFS performs a Breadth-First Search on the graph to find all possible paths
// between the start and end nodes defined in the Stend structure.
func (G *Graph) BFS(Stend StartEnd) {
	var bfs Bfs = Bfs{Visted: make(map[string]bool), Chemin: make(map[int][]string), visitchemin: make(map[string]bool)}

	// If the graph is too large, use DFS instead to avoid memory overload.
	if len(G.Graphs) > 1000 {
		G.Dfs(G.Graphs[Stend.Start].Name, Stend, &bfs)
		delete(bfs.Chemin, bfs.IndexChemin)
		bfs.Reverse()
		bfs.Destribuer()
		return
	}

	queue := make([]string, 0, len(G.Graphs))
	queue = append(queue, G.Graphs[Stend.Start].Name)

	for len(queue) > 0 {
		rom := queue[0]
		queue = queue[1:]
		bfs.Visted[rom] = true

		if rom == G.Graphs[Stend.End].Name {
			bfs.CHemin(G.Graphs[Stend.End].Name, Stend, G)
			continue
		}

		for _, val := range G.Graphs[rom].link {
			if !bfs.Visted[val] {
				queue = append(queue, val)
			}
		}
	}

	delete(bfs.Chemin, bfs.IndexChemin)
	bfs.Reverse()

	for l, val := range bfs.Chemin {
		fmt.Println(l)
		for _, i := range val {
			fmt.Println(i)
		}
	}

	bfs.Destribuer()
}
