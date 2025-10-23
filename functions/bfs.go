package lemino

import (
	"fmt"
	
)

func (G *Graph) BFS(Stend StartEnd) {
	var bfs Bfs = Bfs{Visted: make(map[*Room]bool), Chemin: make(map[int][]*Room), visitchemin: make(map[*Room]bool)}
// var max_chemin int 
	queue := make([]*Room, 0, len(G.Graphs))
	queue = append(queue, G.Graphs[Stend.Start.Name])

	for len(queue) > 0 {
		rom := queue[0]

		queue = queue[1:]

		bfs.Visted[rom] = true
		
		if rom == G.Graphs[Stend.End.Name] {
			
		   bfs.CHemin(G.Graphs[Stend.End.Name], Stend, G)
			
			fmt.Println("OFIGHT")
             continue
		}

		for _, val := range rom.link {
			if !bfs.Visted[val] {
				queue = append(queue, val)
			}
		}
	}
	
	delete(bfs.Chemin,bfs.IndexChemin)
     bfs.Reverse()
	for l, val := range bfs.Chemin {
		fmt.Println(l)
		for _ , i := range val {
			fmt.Println(i)
		}
	}
	
  bfs.Destribuer()

 
}
