package lemino
// BFS performs a Breadth-First Search on the graph to find all possible paths
// from the start room to the end room, and then distributes ants through those paths.
func (G *Graph) BFS(Stend StartEnd) {
	var bfs Bfs = Bfs{Visted: make(map[string]bool), Chemin: make(map[int][]string), visitchemin: make(map[string]bool)}
	if len(G.Graphs) > 1000 {
		G.Dfs(Stend.Start, Stend, &bfs)
		delete(bfs.Chemin, bfs.IndexChemin)
		bfs.Reverse()
		bfs.Destribuer()
		return
	}

	queue := make([]string, 0, len(G.Graphs))
	queue = append(queue, Stend.Start)
	for len(queue) > 0 {

		rom := queue[0]

		queue = queue[1:]

		bfs.Visted[rom] = true

		if rom == Stend.End {
			bfs.CHemin(Stend.End, Stend, G)

			continue
		}

		for _, val := range G.Graphs[rom].link {
			if !bfs.Visted[val] {
				queue = append(queue, val)
				continue
			}
		}
	}

	delete(bfs.Chemin, bfs.IndexChemin)
	bfs.Reverse()

	bfs.Destribuer()
}
