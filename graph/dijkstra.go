package graphStructure

import (
	"math"
)

type Graph struct {
	Nodes   map[string]int
	Edges   map[string]map[string]int
	Visited map[string]bool
	Start   string
	End     string
}

func findLowestCostNode(costs map[string]int, visited map[string]bool) string {
	lowestCost := math.Inf(1)
	lowestCostNode := ""

	for node, cost := range costs {
		if cost < int(lowestCost) && !visited[node] {
			lowestCost = float64(cost)
			lowestCostNode = node
		}
	}

	return lowestCostNode
}

func Dijkstra(g Graph) (map[string]int, map[string]string) {

	costs := make(map[string]int)
	parents := make(map[string]string)

	for node := range g.Edges {
		if node == g.Start {
			costs[node] = 0
		} else {
			costs[node] = int(math.Inf(1))
		}
		parents[node] = ""
	}

	g.Visited = make(map[string]bool)

	node := findLowestCostNode(costs, g.Visited)
	for node != "" {
		cost := costs[node]
		neighbors := g.Edges[node]

		for neighbor, edgeCost := range neighbors {
			newCost := cost + edgeCost
			if newCost < costs[neighbor] {
				costs[neighbor] = newCost
				parents[neighbor] = node
			}
		}

		g.Visited[node] = true

		node = findLowestCostNode(costs, g.Visited)
	}

	return costs, parents
}

func GetShortestPath(parents map[string]string, start, end string) []string {
	path := []string{}
	currentNode := end

	for currentNode != "" {
		path = append([]string{currentNode}, path...)
		currentNode = parents[currentNode]
		if currentNode == start {
			path = append([]string{start}, path...)
			break
		}
	}

	return path
}
