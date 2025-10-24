package lemino

import "fmt"

func (bfs *Bfs) CHemin(rom *Room, Stend StartEnd, G *Graph) bool {

	if rom == G.Graphs[Stend.Start.Name] {
		bfs.Visted[G.Graphs[Stend.Start.Name]] = false
		bfs.Visted[G.Graphs[Stend.End.Name]] = false
	       bfs.Clearmap(G, Stend)
		   bfs.IndexChemin++
		fmt.Println("trouve start apres return d'abord")
		return true
	 }
	 
   bfs.visitchemin[rom] = true
 

for _, va := range rom.link {
		if len(bfs.Chemin[bfs.IndexChemin]) == 0 {
			bfs.Chemin[bfs.IndexChemin] = append(bfs.Chemin[bfs.IndexChemin], G.Graphs[Stend.End.Name])
		}
		// fmt.Println("==>", rom.Name, "links to:")
		// for _, v := range rom.link {
		// 	fmt.Println("   -", v.Name)
		// }

		if _, ok := bfs.Visted[va]; ok {
			if !bfs.visitchemin[va] {
				  bfs.visitchemin[rom] = true
				bfs.Chemin[bfs.IndexChemin] = append(bfs.Chemin[bfs.IndexChemin], va)
				
 	               if bfs.CHemin(va, Stend, G){
                         return true
				   }else {
					bfs.Chemin[bfs.IndexChemin] = bfs.Chemin[bfs.IndexChemin][:len(bfs.Chemin[bfs.IndexChemin])-1]
					bfs.visitchemin[rom] = false
				   }
				
		
			}
		}
	}
	
	return false
	
}
