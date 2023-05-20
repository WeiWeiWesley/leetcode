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

	exampleDFS()
}

func exampleDFS() {

	g := Graph{
		adjList: make(map[*Vertex][]*Edge),
	}

	a := &Vertex{name: "A"}
	b := &Vertex{name: "B"}
	c := &Vertex{name: "C"}
	d := &Vertex{name: "D"}
	e := &Vertex{name: "E"}
	f := &Vertex{name: "F"}

	g.AddVertex(a)
	g.AddVertex(b)
	g.AddVertex(c)
	g.AddVertex(d)
	g.AddVertex(e)
	g.AddVertex(f)

	g.AddEdge(a, b, 1)
	g.AddEdge(a, c, 1)
	g.AddEdge(b, d, 1)
	g.AddEdge(c, f, 1)
	g.AddEdge(d, f, 1)
	g.AddEdge(d, e, 1)

	checkList := map[*Vertex]bool{}

	for k := range g.adjList {
		checkList[k] = false
	}

	fmt.Println("DFS 尋找 a 可到達的點")
	g.DFS(a, checkList)

}

func (g *Graph) DFS(start *Vertex, checkList map[*Vertex]bool) {

	//與其點相鄰邊
	edges := g.adjList[start]

	//把邊上的點印出
	for _, v := range edges {

		if !checkList[v.vertex] {
			fmt.Println(v.vertex.name)
			checkList[v.vertex] = true //避免重複
		}

		g.DFS(v.vertex, checkList) //以新端點向下搜尋
	}
}
