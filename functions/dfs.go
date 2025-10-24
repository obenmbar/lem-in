package lemino

import "fmt"



func (G *Graph) Dfs(rom *Room, Stend StartEnd, BFS *Bfs)  {
    
	if rom == G.Graphs[Stend.End.Name]{
		
		fmt.Println("il trouve un chemin ")
		BFS.CHemin(rom,Stend,G)
		return 
	}
   BFS.Visted[rom]= true
    for _, room := range rom.link {
		if !BFS.Visted[room] {
   G.Dfs(room,Stend,BFS) 
		
	  
		}
	}
}
