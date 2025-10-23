package lemino

func (bfs *Bfs) Clearmap(G *Graph, S StartEnd) {
	if len(bfs.Chemin) >= 2 {
		if len(bfs.Chemin[0]) == 2 && len(bfs.Chemin[1]) == 2 {
			delete(bfs.Chemin, 0)
		}
	}

	for _, i := range bfs.Chemin[bfs.IndexChemin][0].link {
		if len(bfs.Chemin[bfs.IndexChemin]) > 1 {
			if i == bfs.Chemin[bfs.IndexChemin][1] {
				return
			}
		}
	}

	delete(bfs.Chemin, bfs.IndexChemin)
}
