package lemino

// 🔁 CHemin cherche récursivement un chemin entre la salle de fin et celle du début
// Elle enregistre chaque chemin valide dans bfs.Chemin
func (bfs *Bfs) CHemin(rom string, Stend StartEnd, G *Graph) bool {
	if rom == Stend.Start {
		bfs.Visted[Stend.Start] = false
		bfs.Visted[Stend.End] = false
		bfs.Clearmap(G, Stend)
		bfs.IndexChemin++
		return true
	}

	bfs.visitchemin[rom] = true

	for _, va := range G.Graphs[rom].link {
		if len(bfs.Chemin[bfs.IndexChemin]) == 0 {
			bfs.Chemin[bfs.IndexChemin] = append(bfs.Chemin[bfs.IndexChemin], G.Graphs[Stend.End].Name)
		}

		if _, ok := bfs.Visted[va]; ok {
			if !bfs.visitchemin[va] {
				bfs.visitchemin[rom] = true
				bfs.Chemin[bfs.IndexChemin] = append(bfs.Chemin[bfs.IndexChemin], va)

				if bfs.CHemin(va, Stend, G) {
					return true
				} else {
					bfs.Chemin[bfs.IndexChemin] = bfs.Chemin[bfs.IndexChemin][:len(bfs.Chemin[bfs.IndexChemin])-1]
					bfs.visitchemin[rom] = false
				}
			}
		}
	}

	return false
}
