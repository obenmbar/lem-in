package lemino

import "fmt"

func (bfs *Bfs) Destribuer() {
	ants_chemin := make([]int, 0, len(bfs.Chemin))
	TotaleAnts := NumberAnts
	if len(bfs.Chemin) == 1 {
		ants_chemin = append(ants_chemin, NumberAnts)
	} else  {
		for TotaleAnts > 0 {
		
			fmt.Println(len(bfs.Chemin))
			for i := 0; i < len(bfs.Chemin)-1; i++ {
				
				fmt.Println(i)
				if i+1 < len(ants_chemin){
					if (len(bfs.Chemin[i]) + ants_chemin[i]) > (len(bfs.Chemin[i+1]) + ants_chemin[i+1]) {
					fmt.Println("hh")
					TotaleAnts--
					ants_chemin[i+1]++
					break
				} else {
            
					TotaleAnts--
					ants_chemin[i]++
					break
				}
				}
				TotaleAnts--
			}
		}
	} 

	fmt.Println(ants_chemin)
	bfs.Affichage(ants_chemin)
}
