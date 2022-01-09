package breadth_first_search

type Vertex struct {
	Person    Person
	Adjacents []*Vertex
}

type Person struct {
	Name         string
	MusicalGenre string
}
