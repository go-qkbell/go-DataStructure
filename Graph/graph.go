package Graph

// Graph structure
type Graph struct {
	vertices []Vertex
}

type Vertex struct {
	key      int
	adjacent []*Vertex
}
