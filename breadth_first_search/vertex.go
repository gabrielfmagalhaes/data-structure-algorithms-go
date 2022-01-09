package breadth_first_search

type Vertex struct {
	Person     Person
	EnterEdges []Edge
	ExitEdges  []Edge
}

type Person struct {
	Name         string
	MusicalGenre string
}

func (v *Vertex) addEnterEdge(edge Edge) {
	v.EnterEdges = append(v.EnterEdges, edge)
}

func (v *Vertex) addExitEdge(edge Edge) {
	v.ExitEdges = append(v.ExitEdges, edge)
}
