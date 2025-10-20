package lemino

type Room struct {
	Name string
	X    int
	Y    int
	link []*Room
}
type Graph struct {
	Graphs map[string]*Room
}
type StartEnd struct {
	Start Room
	End  Room
}
type Bfs struct {
	Visted map[*Room]bool
	Chemin map[int][]*Room
	i int
	path [][]*Room
}
var (
	err error
)

