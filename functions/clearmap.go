package lemino

func (bfs *Bfs) Clearmap(G *Graph, S StartEnd){
		   for _, i := range bfs.Chemin[bfs.IndexChemin][0].link {
			if len(bfs.Chemin[bfs.IndexChemin]) > 1 {
				   if i == bfs.Chemin[bfs.IndexChemin][1] {
				return
			  }
		   }
			}
		
          delete(bfs.Chemin, bfs.IndexChemin)
	}

