package lemino

import "fmt"
// AddLinks adds a bidirectional link between two rooms in the graph.
// It performs several validations to ensure that both rooms exist, are distinct,
// and are not already linked together before creating the connection.
func (G *Graph) AddLinks(link []string, stend StartEnd) error {
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
			if key := val; key == link[1] {
				return fmt.Errorf("Error en add link , les deux room deja liee avec eux")
			}
		}
	}

	G.Graphs[link[0]].link = append(G.Graphs[link[0]].link, link[1])
	G.Graphs[link[1]].link = append(G.Graphs[link[1]].link, link[0])

	return nil
}
