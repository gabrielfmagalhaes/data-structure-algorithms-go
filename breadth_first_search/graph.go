package breadth_first_search

type Graph struct {
	Vertices []Vertex
	Edges    []Edge
}

func (g *Graph) AddVertex(person Person) {
	vertex := Vertex{Person: person}

	g.Vertices = append(g.Vertices, vertex)
}

func (g *Graph) AddEdge(personVisited, personNeighbor Person) {
	indexVisited := g.getVertexPos(personVisited)
	indexNeighbor := g.getVertexPos(personNeighbor)

	edge := Edge{g.Vertices[indexVisited], g.Vertices[indexNeighbor]}

	g.Vertices[indexVisited].addEnterEdge(edge)
	g.Vertices[indexVisited].addExitEdge(edge)

	g.Edges = append(g.Edges, edge)
}

func (g *Graph) BreadthFirst() Person {
	var visitedVertices []Vertex
	var queueVertices []Vertex

	queueVertices = append(queueVertices, g.Vertices[0])

	for len(queueVertices) > 0 {
		visited := queueVertices[0]

		for i := 0; i < len(visited.ExitEdges); i++ {
			neighbor := visited.ExitEdges[i].ToVertex

			if doesHeEnjoyRock(neighbor.Person) {
				return neighbor.Person
			}

			if !wasVertexAlreadyVisited(neighbor, visitedVertices) {
				pos := g.getVertexPos(neighbor.Person)

				queueVertices = append(queueVertices, g.Vertices[pos])
				visitedVertices = append(visitedVertices, neighbor)
			}
		}
		queueVertices = queueVertices[1:]
	}

	return Person{}
}

func (g *Graph) getVertexPos(person Person) int {
	for index, elem := range g.Vertices {
		if elem.Person.Name == person.Name {
			return index
		}
	}

	return 0
}

func wasVertexAlreadyVisited(vertex Vertex, visitedVertices []Vertex) bool {
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
