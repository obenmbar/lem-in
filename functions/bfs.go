package lemino

import "fmt"


func (G *Graph) EdmondKarp(Stend StartEnd) {
	var allPaths [][]string
    var bfs *Bfs

	for {
	
		bfs = &Bfs{
			Visted: make(map[string]bool),
			Chemin: make(map[int][]string),
			Parent: make(map[string]string),
		}

		 G.BFSOnce(Stend, bfs)
		if bfs.Chemin[bfs.IndexChemin] == nil {
			break // ما بقاوش طرق
		}
		fmt.Println("a",bfs.Chemin)
		
	
		G.removeEdges(bfs.Chemin[bfs.IndexChemin]) // نحيد الروابط ديال الطريق باش نلقاو طريق جديد
	}

	fmt.Println("جميع الطرق الممكنة:", allPaths)
	fmt.Println(len(allPaths))
	fmt.Println(len(bfs.Chemin))
	fmt.Println("hhh",bfs.IndexChemin)
}

func (G *Graph) removeEdges(path []string) {
	for i := 0; i < len(path)-1; i++ {
		a, b := path[i], path[i+1]
		G.removeLink(a, b)
	}
}

// دالة مساعدة صغيرة
func (G *Graph) removeLink(a, b string) {
	links := G.Graphs[a].link
	newLinks := []string{}
	for _, v := range links {
		if v != b {
			newLinks = append(newLinks, v)
		}
	}
	G.Graphs[a].link = newLinks
}

func (G *Graph) BFSOnce(Stend StartEnd, bfs *Bfs)  {
	queue := []string{Stend.Start}

	for len(queue) > 0 {
		rom := queue[0]
		queue = queue[1:]

		if rom == Stend.End {
			bfs.Allchemin(Stend)
			bfs.IndexChemin++
			
		}

		for _, val := range G.Graphs[rom].link {
			if !bfs.Visted[val] {
				bfs.Visted[val] = true
				bfs.Parent[val] = rom
				queue = append(queue, val)
			}
		}
	}
	
}
