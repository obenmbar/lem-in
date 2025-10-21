package lemino

import "fmt"

func (G *Graph) BFS() {
	var (
		bfs   Bfs = Bfs{Visted: make(map[*Room]bool), Chemin: make(map[int][]*Room)}
		stend StartEnd
	)
	queue := make([]*Room, 0, len(G.Graphs))
	queue = append(queue, &stend.Start)

	for len(queue) > 0 {
		rom := queue[0]
		queue = queue[1:]
		bfs.Visted[rom] = true
		if rom.Name == stend.End.Name && rom.X == stend.End.X && rom.Y == stend.Start.Y {
			for key, _ := range bfs.Visted {
				bfs.Chemin[bfs.i]= 	append(bfs.Chemin[bfs.i], key)
			}
              bfs.i++
			fmt.Println("ofight")
			break
		}
		for _, val := range rom.link {
			if !bfs.Visted[val] {
				queue = append(queue, val)
			}
		}
	}
	bfs.path = make([][]*Room,0) 
     for _, val := range stend.End.link {
		for key, value := range bfs.Chemin{
			for _, romm := range value{
               if romm == val {
               bfs.path[key] = append(bfs.path[key], romm)
			   }
			}
		}
	 }
	 fmt.Println(bfs.path)
	fmt.Println("ok")
}
