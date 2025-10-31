package lemino

import (
	"fmt"
	"strconv"
	"strings"
)

// Parser analyzes and validates the input data describing the ant farm (lem-in).
// It extracts the number of ants, rooms, start and end points, and tunnels between rooms,
// then builds the graph and launches the BFS algorithm.
func Parser(Data string) {
	var (
		R     Room
		Gra   Graph = Graph{make(map[string]*Room),make(map[string]map[string]int)}
		Stend StartEnd

		IsStart             bool
		IsEnd               bool
		IsTunnel            bool
		IsNumberAnts        bool
		checkStartEndNumber int
		Index_start         int
	)

	// Split the input data into lines
	SliceData := strings.Split(Data, "\n")

	for i, valu := range SliceData {
		value := strings.TrimSpace(valu)

		// Handle comments and special commands like ##start / ##end
		if strings.HasPrefix(value, "#") {
			if value == "##start" {
				// Define the start room
				if !IsStart {
					for v := i + 1; v < len(SliceData); v++ {
						value_afterstart := strings.TrimSpace(SliceData[v])
						if value_afterstart == "" || strings.HasPrefix(value_afterstart, "#") {
							if value_afterstart != "##end" {
								Index_start++
								continue
							} else {
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
					Stend.Start = R.Name
					IsStart = true
					checkStartEndNumber++
					continue

				} else {
					checkStartEndNumber++
					continue
				}

			} else if value == "##end" {
				// Define the end room
				if !IsEnd {
					for v := i + 1; v < len(SliceData); v++ {
						value_after_end := strings.TrimSpace(SliceData[v])
						if value_after_end == "" || strings.HasPrefix(value_after_end, "#") {
							if strings.TrimSpace(SliceData[v]) != "##start" {
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
					Stend.End = R.Name
					continue
				} else {
					checkStartEndNumber++
					continue
				}
			} else {
				continue
			}

		} else if value == "" {
			continue

		// Detect the number of ants (must be before rooms)
		} else if !IsStart && !IsNumberAnts {
			NumberAnts, err = strconv.Atoi(value)
			if err != nil || NumberAnts <= 0 {
				fmt.Println("le nombre de fourmis n'est pas valide")
				return
			}
			IsNumberAnts = true
			continue

		// Detect and add rooms
		} else if len(value) >= 5 && strings.Count(value, " ") > 0 {
			if IsTunnel {
				fmt.Println("error en l'ordre du FORMAT DE TEXT")
				return
			}
			err = R.AddRoom(value)
			if err != nil {
				fmt.Println(err)
				return
			}
			err = Gra.Addmap(R)
			if err != nil {
				fmt.Println(err)
				return
			}
			continue

		// Detect and add tunnels (links)
		} else if strings.Count(value, "-") == 1 {
			if strings.Count(value, " ") == 0 && len(value) >= 3 {
				IsTunnel = true
				link := strings.Split(value, "-")

				err = Gra.AddLinks(link, Stend)
				if err != nil {
					fmt.Println(err)
					return
				}

			} else {
				fmt.Println("error en format de tunele exemple *4-3*")
				return
			}

		// Handle invalid lines
		} else {
			fmt.Println("error en FORMAT des donnees de text.txt : invalid syntax")
			return
		}
	}

	// Validate that both start and end exist exactly once
	if checkStartEndNumber != 2 {
		if !IsStart {
			fmt.Println("error : missing ##start in text.txt")
			return
		} else if !IsEnd {
			fmt.Println("error : missing ##end in text.txt")
			return
		} else {
			fmt.Println("error : duplicated ##start or ##end")
			return
		}
	}

	// Validate the number of ants was provided
	if !IsNumberAnts {
		fmt.Println("error : missing ant count at the beginning")
		return
	}

	// Launch BFS algorithm to find paths

	Gra.BFS(Stend)
	
}
