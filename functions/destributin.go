package lemino

import "fmt"

func (bfs *Bfs) Destribuer() {
	ants_chemin := make([]int, len(bfs.Chemin))
	TotaleAnts := NumberAnts
	if len(bfs.Chemin) == 1 {
		ants_chemin[0] = NumberAnts
	} else {
		for TotaleAnts > 0 {
			for i := 0; i < len(bfs.Chemin)-1; i++ {
				if TotaleAnts == 0 {
					break
				}

				if (len(bfs.Chemin[i]) + ants_chemin[i]) > (len(bfs.Chemin[i+1]) + ants_chemin[i+1]) {
					TotaleAnts--
					ants_chemin[i+1]++

				} else {
					TotaleAnts--
					ants_chemin[i]++

				}
			}
		}
	}

	fmt.Println(ants_chemin)
	bfs.Affichage(ants_chemin)
}
