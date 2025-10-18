package lemino

func (G *Graph) Addmap(R *Room) {
	if _, ok := G.Graphs[R.Name]; !ok {
		G.Graphs[R.Name] = R
	}
}
