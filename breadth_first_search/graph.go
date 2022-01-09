package breadth_first_search

import "fmt"

type Graph struct {
	Vertices []*Vertex
}

func (g *Graph) AddVertex(person Person) {
	if contains(g.Vertices, person) {
		err := fmt.Errorf("Vertex %v not added because it is an existing key", person)
		fmt.Println(err.Error())
	} else {
		g.Vertices = append(g.Vertices, &Vertex{Person: person})
	}
}

func (g *Graph) AddEdge(personVisited, personNeighbor Person) {
	fromVertex := g.getVertex(personVisited)
	toVertex := g.getVertex(personNeighbor)

	if fromVertex == nil || toVertex == nil {
		err := fmt.Errorf("Invalid edge (%v --> %v)", personVisited, personNeighbor)
		fmt.Println(err.Error())
	} else if contains(fromVertex.Adjacents, personNeighbor) {
		err := fmt.Errorf("Existing edge (%v --> %v)", personVisited, personNeighbor)
		fmt.Println(err.Error())
	} else {
		fromVertex.Adjacents = append(fromVertex.Adjacents, toVertex)
	}
}

func (g *Graph) getVertex(person Person) *Vertex {
	for i, elem := range g.Vertices {
		if elem.Person.Name == person.Name {
			return g.Vertices[i]
		}
	}

	return nil
}

func contains(v []*Vertex, person Person) bool {
	for _, elem := range v {
		if elem.Person.Name == person.Name {
			return true
		}
	}

	return false
}

func (g *Graph) BreadthFirst() Person {
	var visitedVertices []*Vertex
	var queueVertices []*Vertex

	queueVertices = append(queueVertices, g.Vertices[0])

	for len(queueVertices) > 0 {
		visited := queueVertices[0]

		for i := 0; i < len(visited.Adjacents); i++ {
			neighbor := visited.Adjacents[i]

			if doesHeEnjoyRock(neighbor.Person) {
				return neighbor.Person
			}

			if !wasVertexAlreadyVisited(neighbor, visitedVertices) {
				findVertex := g.getVertex(neighbor.Person)

				queueVertices = append(queueVertices, findVertex)
				visitedVertices = append(visitedVertices, neighbor)
			}
		}
		queueVertices = queueVertices[1:]
	}

	return Person{}
}

func wasVertexAlreadyVisited(vertex *Vertex, visitedVertices []*Vertex) bool {
	for _, elem := range visitedVertices {
		if vertex.Person.Name == elem.Person.Name {
			return true
		}
	}

	return false
}

func doesHeEnjoyRock(person Person) bool {
	return person.MusicalGenre == "Rock"
}
