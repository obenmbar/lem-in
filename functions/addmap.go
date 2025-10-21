package lemino

import "fmt"

func (G *Graph) Addmap(R Room) error {
	if _, ok := G.Graphs[R.Name]; !ok {
		for _, value := range G.Graphs {
			if value.X == R.X && value.Y == R.Y {
				return fmt.Errorf("error en add map, les cordoner de cette Room deja existe dans une autre room ")
			}
		}
		G.Graphs[R.Name] = &R
	} else {
			return fmt.Errorf("error en add map, cette room deja exist")

	}
	return nil
}
