package lemino

import "fmt"

func (bfs *Bfs) Allchemin(stend StartEnd) {
	for room := stend.End; room != ""; room = bfs.Parent[room] {
		bfs.Chemin[bfs.IndexChemin] = append([]string{room}, bfs.Chemin[bfs.IndexChemin]...)
		if room == stend.Start {
			break
		}

	}
	fmt.Println("hihi",bfs.IndexChemin)
	fmt.Println("hhhhhhhh", bfs.Chemin[bfs.IndexChemin])
}
