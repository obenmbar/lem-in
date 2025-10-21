package lemino

import (
	"fmt"
)

func (bfs *Bfs) CHemin(rom *Room, Stend StartEnd, G *Graph) {
	if rom == G.Graphs[Stend.Start.Name] {
		bfs.Clearmap(G, Stend)
		bfs.IndexChemin++
		return
	}

	if rom != G.Graphs[Stend.End.Name] || rom != G.Graphs[Stend.Start.Name] {
		bfs.visitchemin[rom] = true
	}

	for _, va := range rom.link {
		if len(bfs.Chemin[bfs.IndexChemin]) == 0 {
			bfs.Chemin[bfs.IndexChemin] = append(bfs.Chemin[bfs.IndexChemin], G.Graphs[Stend.End.Name])
		}
		if _, ok := bfs.Visted[va]; ok {
			if !bfs.visitchemin[va] {

				bfs.Chemin[bfs.IndexChemin] = append(bfs.Chemin[bfs.IndexChemin], va)
				fmt.Println(va)
				bfs.CHemin(va, Stend, G)
			}
		}
	}
}
