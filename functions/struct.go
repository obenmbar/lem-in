package lemino

type Room struct {
	Name string
	X    int
	Y    int
	link []string
}
type Graph struct {
	Graphs map[string]*Room
	Capacitie map[string]map[string]int
	
}
type StartEnd struct {
	Start string
	End   string
}
type Bfs struct {
	Visted      map[string]bool
	Chemin      map[int][]string
	visitchemin map[string]bool
	Parent map[string]string
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
 
 type Edge struct {
	From string
	To string
	Capacity int 
	Reverse *Edge
 }