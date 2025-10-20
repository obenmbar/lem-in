package lemino

import "fmt"

func (G *Graph) AddLinks(link []string) error {
	if len(link) != 2 {
		return fmt.Errorf("error en add link ,le number de room n'est pas valable")
	}
	if link[0] == link[1] {
		return fmt.Errorf("error en add link , les deux room sont identique")
	}
	if link[0] == "" || link[1] == "" {
		return fmt.Errorf("error en add link , le nom de room est vide")
	}

	if _, ok := G.Graphs[link[0]]; !ok {
		return fmt.Errorf("error en add link , le premier link n'est pas exist dans le map")
	} else if _, ok = G.Graphs[link[1]]; !ok {
		return fmt.Errorf("error en add link , le deuxieme  link n'est pas exist dans le map")
	}

	if len(G.Graphs[link[0]].link) > 0 {
		for _, val := range G.Graphs[link[0]].link {
			if key := G.Graphs[val.Name]; key.Name == link[1]{
				return fmt.Errorf("Error en add link , les deux room deja liee avec eux")
			}
		}
	}

	G.Graphs[link[0]].link = append(G.Graphs[link[0]].link, G.Graphs[link[1]])
	G.Graphs[link[1]].link = append(G.Graphs[link[1]].link, G.Graphs[link[0]])

	return nil
}
