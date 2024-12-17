package dataStructure

import "fmt"

type EdgeStruct struct {
	Value     string
	Neighbors []string
}

type GraphStruct struct {
	nodes map[string]EdgeStruct
}

func (g *GraphStruct) Push(node string, neighbors []string) {
	if _, exists := g.nodes[node]; !exists {
		g.nodes[node] = EdgeStruct{Value: node, Neighbors: neighbors}
	}
}

func (g *GraphStruct) PushNeighbors(node string, neighbors []string) {
	if _, exists := g.nodes[node]; exists {
		current := g.nodes[node]
		current.Neighbors = append(current.Neighbors, neighbors...)
		g.nodes[node] = current
	}

}

func (g GraphStruct) Search(node string) EdgeStruct {
	queue := Queue()
	queue.push(node)
	verifies := make(map[string]bool)

	for queue.size > 0 {
		queueItem := queue.pop()

		if !verifies[queueItem] {
			verifies[queueItem] = true

			if node, exists := g.nodes[queueItem]; exists {
				fmt.Printf("Visitando nó: %s, Fila atual: %v\n", queueItem, queue.items)
				for _, neighbor := range node.Neighbors {
					if !verifies[neighbor] {
						queue.push(neighbor)
						fmt.Printf("Visitando vizinhos: %s, Fila atual: %v\n", queueItem, queue.items)
					}
				}
				return node
			}
		}
	}

	return EdgeStruct{}
}

func NewGraph() *GraphStruct {
	return &GraphStruct{nodes: make(map[string]EdgeStruct)}
}
