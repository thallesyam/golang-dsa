package main

import (
	"example/dsa/dataStructure"
	graphStructure "example/dsa/graph"
	"example/dsa/order"
	"fmt"
)

func main() {
	quicksortArr := []int{3, 5, 2, 1, 4}

	fmt.Printf("Quicksort: ")
	fmt.Println(order.Quicksort(quicksortArr))

	fmt.Printf("HashMap: ")
	dataStructure.HashMapMain()

	fmt.Println("Graph")
	graph := dataStructure.NewGraph()
	graph.Push("A", []string{"B", "C"})
	graph.Push("B", []string{"D", "E"})
	graph.Push("C", []string{"F", "G"})
	graph.Push("D", []string{})
	graph.Push("E", []string{"H"})
	graph.Push("F", []string{})
	graph.Push("G", []string{})
	graph.Push("H", []string{})

	// Cenário 1: Testando busca por um nó no grafo
	fmt.Println("\nBuscando o nó 'E'")
	node := graph.Search("E")
	if node.Value != "" {
		fmt.Printf("Nó encontrado: %s com vizinhos: %v\n", node.Value, node.Neighbors)
	} else {
		fmt.Println("Nó não encontrado.")
	}

	g := graphStructure.Graph{
		Nodes: map[string]int{},
		Edges: map[string]map[string]int{
			"A": {"B": 2, "C": 5},
			"B": {"C": 1, "D": 4},
			"C": {"D": 2},
			"D": {},
		},
		Start: "A",
		End:   "D",
	}

	costs, parents := graphStructure.Dijkstra(g)
	path := graphStructure.GetShortestPath(parents, g.Start, g.End)

	fmt.Println("Costs:", costs)
	fmt.Println("Shortest Path:", path)

}
