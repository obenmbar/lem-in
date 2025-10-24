package lemino

func (bfs *Bfs) Clearmap(G *Graph, S StartEnd) {
	// fmt.Println("kkkkkkkkkkk")
	// fmt.Println(len(bfs.Chemin[bfs.IndexChemin]))
	// if len(bfs.Chemin[bfs.IndexChemin]) < 2 {
	// 	delete(bfs.Chemin, bfs.IndexChemin)
	// }
	for _, i := range bfs.Chemin[bfs.IndexChemin][0].link {
		if len(bfs.Chemin[bfs.IndexChemin]) > 1 {
			if i == bfs.Chemin[bfs.IndexChemin][1] {
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
