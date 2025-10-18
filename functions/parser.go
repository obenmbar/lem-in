package lemino

import (
	"fmt"
	"strconv"
	"strings"
)

func Parser(Data string) {
	var (
		R     Room
		Stend StartEnd = StartEnd{Start: &R, End: &R}
		Gra   Graph    = Graph{make(map[string]*Room)}

		NumberAnts          int
		IsStart             bool
		IsEnd               bool
		IsNumberAnts        bool
		checkStartEndNumber int
	)

	SliceData := strings.Split(Data, "\n")
	for i, value := range SliceData {
		if strings.TrimSpace(value) == "" {
			continue
		} else if !IsStart && !IsNumberAnts {
			NumberAnts, err = strconv.Atoi(value)
			if err != nil || NumberAnts <= 0 {
				fmt.Println(" le nombre de fourmis n'est pas valid")
				return
			}
			IsNumberAnts = true
			continue
		}

		if strings.HasPrefix(value, "#") {
			if value == "##start" {
				if !IsStart {
					var Index_start int
					if strings.TrimSpace(SliceData[i+1]) == "" {
						Index_start += 1
						continue
					} else {
						err = R.AddRoom(SliceData[i+Index_start+1])
						if err != nil {
							fmt.Println("55")
							fmt.Println(err)
							return
						}
						Stend.Start.Name = R.Name
						Stend.Start.X = R.X
						Stend.Start.Y = R.Y
						Gra.Addmap(&R)
						IsStart = true
						checkStartEndNumber++
						continue
					}

				} else {
					checkStartEndNumber++
					continue
				}
			} else if value == "##end" {
				if !IsEnd {
					IsEnd = true
					checkStartEndNumber++
					err = R.AddRoom(SliceData[i+1])
					if err != nil {
						fmt.Println(err)
						return
					}
					Stend.End.Name = R.Name
					Stend.End.X = R.X
					Stend.End.Y = R.Y
					Gra.Addmap(&R)
					continue
				} else {
					checkStartEndNumber++
					continue
				}
			} else {
				continue
			}
		} else if len(value) >= 5 && strings.Count(value, " ") > 0 {

			err = R.AddRoom(value)
			if err != nil {
				fmt.Println(err)
				return
			}

			Gra.Addmap(&R)
			continue
		} else if strings.Count(value, "-") == 1 {
			if strings.Count(value, " ") != 0 {
				link := strings.Split(value, "-")
				err = Gra.AddLinks(link)
				if err != nil {
					fmt.Println(err)
					return
				}

			}
		} else {
			fmt.Println("error en FORMAT  des donnees de text.txt il y'a un sembole n'exist pas com un seulement chaine des caracter simple ")
			return
		}
	}
	if checkStartEndNumber != 2 {
		if !IsStart {
			fmt.Println("error en FORMAT de test.txt, il n'ya pas la valeur '##start'")
			return
		} else if !IsEnd {
			fmt.Println("error en FORMAT de test.txt, il n'ya pas la valeur '##End'")
			return
		} else {
			fmt.Println("error en FORMAT de test.txt, il y a un repetetion de *##start* ou*##end*")
			return
		}
	}
	if !IsNumberAnts {
		fmt.Println("error , il n ya pas un valeur senifie le nombre des fourmis ")
		return
	}
	fmt.Println("ok")
}
