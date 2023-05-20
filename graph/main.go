package main

import (
	"fmt"
	"strings"
)

type Vertex struct {
	name string
}

type Edge struct {
	vertex   *Vertex
	distance int
}

type Graph struct {
	adjList map[*Vertex][]*Edge
}

func (g *Graph) AddVertex(vertex *Vertex) {
	g.adjList[vertex] = []*Edge{}
}

func (g *Graph) AddEdge(a, b *Vertex, distiance int) {
	g.adjList[a] = append(g.adjList[a], &Edge{
		vertex:   b,
		distance: distiance,
	})
}

func (g *Graph) Show() {

	for k, v := range g.adjList {

		edge := []string{}

		if v != nil {
			for i := range v {
				edge = append(edge, v[i].vertex.name)
			}
		}

		fmt.Printf("[%s] %s\n", k.name, strings.Join(edge, ","))
	}

}

func main() {

	g := Graph{
		adjList: make(map[*Vertex][]*Edge),
	}

	a := &Vertex{name: "A"}
	b := &Vertex{name: "B"}
	c := &Vertex{name: "C"}

	g.AddVertex(a)
	g.AddVertex(b)
	g.AddVertex(c)

	g.AddEdge(a, b, 2)
	g.AddEdge(a, c, 3)
	g.AddEdge(b, c, 1)

	g.Show()
}
