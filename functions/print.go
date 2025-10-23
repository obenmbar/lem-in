package lemino

import "fmt"

func (bfs *Bfs) Affichage(ants_chemin []int) {
	var SliceAnt []*Ant
	antId := 1

	// --- إنشاء النمل حسب عددهم والمسارات ---
	for index, count := range ants_chemin {
		for i := 0; i < count; i++ {
			SliceAnt = append(SliceAnt, &Ant{
				Id:   antId,
				Path: bfs.Chemin[index],
				Step: -1, // -1 يعني باقي ما بداش
			})
			antId++
		}
	}

	totalAnts := len(SliceAnt)
	finished := 0
	time := 0 // عدد الجولات

	// --- المحاكاة ---
	for finished < totalAnts {
		line := ""

		for i := 0; i < len(SliceAnt); i++ {
			ant := SliceAnt[i]

			// النمل يدخل تدريجياً: كل جولة، نملة جديدة تبدأ طريقها
			if ant.Step == -1 {
				if i <= time { // واحد جديد يدخل كل جولة
					ant.Step = 0
				} else {
					continue
				}
			}

			// إذا سالات الطريق، ما ندير والو
			if ant.Isfiniched {
				continue
			}

			// الخطوة القادمة
			ant.Step++
			if ant.Step < len(ant.Path) {
				line += fmt.Sprintf("L%d-%s ", ant.Id, ant.Path[ant.Step].Name)
			}

			// تحقق واش وصل للنهاية
			if ant.Step == len(ant.Path)-1 {
				ant.Isfiniched = true
				finished++
			}
		}

		// طبع الخط إذا كاين تحرك
		if line != "" {
			fmt.Println(line)
		}

		time++ // الجولة القادمة
	}
}
