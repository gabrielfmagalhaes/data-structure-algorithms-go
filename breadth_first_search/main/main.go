package main

import (
	"fmt"
	graph "grokking-algorithms-go/breadth_first_search"
)

func main() {
	newGraph := graph.Graph{}

	gabriel := graph.Person{Name: "Gabriel", MusicalGenre: "Samba"}
	john := graph.Person{Name: "John", MusicalGenre: "Blues"}
	marie := graph.Person{Name: "Marie", MusicalGenre: "Classical"}
	diego := graph.Person{Name: "Diego", MusicalGenre: "Country"}
	leonardo := graph.Person{Name: "Leonardo", MusicalGenre: "Funk"}
	rafael := graph.Person{Name: "Rafael", MusicalGenre: "MPB"}
	lucas := graph.Person{Name: "Lucas", MusicalGenre: "Rock"}
	daniella := graph.Person{Name: "Daniella", MusicalGenre: "Funk"}

	newGraph.AddVertex(gabriel)
	newGraph.AddVertex(john)
	newGraph.AddVertex(marie)
	newGraph.AddVertex(diego)
	newGraph.AddVertex(leonardo)
	newGraph.AddVertex(rafael)
	newGraph.AddVertex(lucas)
	newGraph.AddVertex(daniella)

	newGraph.AddEdge(gabriel, john)
	newGraph.AddEdge(gabriel, diego)
	newGraph.AddEdge(gabriel, marie)

	newGraph.AddEdge(john, leonardo)
	newGraph.AddEdge(john, rafael)

	newGraph.AddEdge(marie, leonardo)
	newGraph.AddEdge(marie, lucas)
	newGraph.AddEdge(marie, daniella)

	fmt.Println(newGraph.BreadthFirst())
}
