package lemino

type Room struct {
	Name string
	X    int
	Y    int
	link []string
}
type Graph struct {
	Graphs map[string]*Room
}
type StartEnd struct {
	Start string
	End   string
}
type Bfs struct {
	Visted      map[string]bool
	Chemin      map[int][]string
	visitchemin map[string]bool
	
	IndexChemin int
}

var (
	NumberAnts int
	err        error

)
 type Ant struct {
	Id int 
	Path []string
	Isfiniched bool
	 Step int
	
 }