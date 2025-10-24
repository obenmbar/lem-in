package lemino

import "fmt"

func (bfs *Bfs) Affichage(ants_chemin []int) {

	var SliceAnt []*Ant
	var i int = 1
	for i <= NumberAnts {

		for index := 0; index < len(ants_chemin); index++ {
			if ants_chemin[index] > 0 {
				i = i + index
				SliceAnt = append(SliceAnt, &Ant{
					Id:   i,
					Path: bfs.Chemin[index],
					Step: 0,
				})
			
				ants_chemin[index]--
			}
		}
		i++
	}

	totalAnts := len(SliceAnt)
	finished := 0
	for finished < totalAnts {
		line := ""

		room_ocupe := make(map[string]bool)
		for i := 0; i < len(SliceAnt); i++ {

			formi := SliceAnt[i]

			if formi.Step == 0 {
				formi.Step++
			}

		
			
			if room_ocupe[formi.Path[formi.Step].Name] {
		
				continue
			}


			if formi.Isfiniched {
				continue
			}

			if formi.Step < len(formi.Path) {
				if formi.Step != len(formi.Path)-1 {
					room_ocupe[formi.Path[formi.Step].Name] = true
				}

				line += fmt.Sprintf("L%d-%s ", formi.Id, formi.Path[formi.Step].Name)

			}

			if formi.Step == len(formi.Path)-1 {
				finished++
				formi.Isfiniched = true
				continue
			}
			formi.Step++
		}
		if line != "" {
			fmt.Println(line)
		}

	}
}
