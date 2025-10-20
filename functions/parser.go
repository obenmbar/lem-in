package lemino

import (
	"fmt"
	"strconv"
	"strings"
)

func Parser(Data string) {
	var (
		R     Room
		Stend StartEnd 
		Gra   Graph    = Graph{make(map[string]*Room)}

		NumberAnts          int
		IsStart             bool
		IsEnd               bool
		IsNumberAnts        bool
		checkStartEndNumber int
		Index_start         int
	)

	SliceData := strings.Split(Data, "\n")
	for i, value := range SliceData {
		if strings.HasPrefix(strings.TrimSpace(value), "#") {
			if strings.TrimSpace(value) == "##start" {
				if !IsStart {
					for v := i + 1; v < len(SliceData); v++ {
						if strings.TrimSpace(SliceData[v]) == ""  || strings.HasPrefix(strings.TrimSpace(SliceData[v]),"#"){
							if strings.TrimSpace(SliceData[v]) != "##end" {
								Index_start += 1
								
							     continue
							}else {
								fmt.Println("il ya end dessous de start alors les cordonne pas exist")
								return
							}
							
						} else {
							break
						}
					}

					err = R.AddRoom(SliceData[i+Index_start+1])
					Index_start = 0
					if err != nil {
						fmt.Println(err)
						return
					}
					Stend.Start = R
					IsStart = true
					checkStartEndNumber++
					continue

				} else {
					checkStartEndNumber++
					continue
				}
			} else if strings.TrimSpace(value) == "##end" {
				if !IsEnd {
					for v := i + 1; v < len(SliceData); v++ {
						if strings.TrimSpace(SliceData[v]) == "" || strings.HasPrefix(SliceData[v], "#")  {
							if  strings.TrimSpace(SliceData[v]) != "##start" {
								Index_start += 1
							     continue
							} else {
								fmt.Println("error il ya la valeur ##start directement desous de la valeur ##end")
								return
							}
						} else {
							break
						}
					}
					IsEnd = true
					checkStartEndNumber++
					err = R.AddRoom(SliceData[i+Index_start+1])
					Index_start = 0
					if err != nil {
						fmt.Println(err)
						return
					}
					Stend.End = R
					continue
				} else {
					checkStartEndNumber++
					continue
				}
			} else {
				continue
			}
		} else if strings.TrimSpace(value) == "" {
			continue
		} else if !IsStart && !IsNumberAnts {
			NumberAnts, err = strconv.Atoi(strings.TrimSpace(value))
			if err != nil || NumberAnts <= 0 {
				fmt.Println(" le nombre de fourmis n'est pas valid")
				return
			}
			IsNumberAnts = true
			continue
		} else if len(strings.TrimSpace(value)) >= 5 && strings.Count(value, " ") > 0 && strings.Count(value,"-")==0 {

			err = R.AddRoom(value)
			if err != nil {
		
				fmt.Println(err)
				return
			}
			err = Gra.Addmap(R)
			if err != nil {
				fmt.Println(R)
				
				fmt.Println(err)
				return
			}
			continue
		} else if strings.Count(value, "-") == 1 {
			if strings.Count(strings.TrimSpace(value), " ") == 0 && len(value) >= 3 {
				link := strings.Split(strings.TrimSpace(value), "-")
                 
				err = Gra.AddLinks(link)
				if err != nil {
				
					fmt.Println(err)
					return
				}
			} else {
				fmt.Println("error en format de tunele exemple *4-3*")
				return
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
	Gra.BFS()
}
