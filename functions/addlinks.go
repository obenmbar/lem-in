package lemino

import "fmt"

func (G *Graph) AddLinks(link []string) error {
	if len(link) != 2 {
		return fmt.Errorf("error en add link ,le number de room n'est pas valable")
	}
	if _, ok := G.Graphs[link[0]]; !ok {
		return fmt.Errorf("error en add link , le premier link n'est pas exist dans le map")
	} else if _, ok = G.Graphs[link[1]]; !ok {
		return fmt.Errorf("error en add link , le deuxieme  link n'est pas exist dans le map")
	}
	G.Graphs[link[0]].link = append(G.Graphs[link[0]].link, G.Graphs[link[1]])
	G.Graphs[link[0]].link = append(G.Graphs[link[1]].link, G.Graphs[link[0]])
	return nil 
}
