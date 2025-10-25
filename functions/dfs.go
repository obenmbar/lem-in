package lemino
// Dfs performs a Depth-First Search on the graph starting from the given room,
// recursively exploring all possible paths until the end room is found.
func (G *Graph) Dfs(rom string, Stend StartEnd, BFS *Bfs) {
	if rom == Stend.End {
		BFS.CHemin(rom, Stend, G)
		return
	}
	BFS.Visted[rom] = true
	for _, room := range G.Graphs[rom].link {
		if !BFS.Visted[room] {
			G.Dfs(room, Stend, BFS)
		}
	}
}
