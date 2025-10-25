package lemino
// Clearmap removes the current invalid or blocked path from the BFS results
// and reindexes the remaining valid paths to keep the path order consistent.
func (bfs *Bfs) Clearmap(G *Graph, S StartEnd) {

	for _, chombre := range G.Graphs[bfs.Chemin[bfs.IndexChemin][0]].link {
		if len(bfs.Chemin[bfs.IndexChemin]) > 1 {
			if chombre == bfs.Chemin[bfs.IndexChemin][1] {
				return
			}
		}
	}

	delete(bfs.Chemin, bfs.IndexChemin)
	index := 0
	for key, value := range bfs.Chemin {
		if key != index {
			bfs.Chemin[index] = value
			delete(bfs.Chemin, key)
		
			index++
			continue
		}
		index++
	}
}
