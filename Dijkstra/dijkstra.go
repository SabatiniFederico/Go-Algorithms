package dijkstra

import (
	"math"

	"github.com/x1m3/priorityQueue"
)

type Node struct {
	name   string
	weight int
}

// I need to define what priority means in a struct
func (i Node) HigherPriorityThan(other priorityQueue.Interface) bool {
	return i.weight < other.(Node).weight
}

func dijkstra(graph map[string][]Node, root string) map[string]int {

	distances := make(map[string]int)

	for key := range graph {
		distances[key] = math.MaxInt32
	}
	distances[root] = 0

	queue := priorityQueue.New()
	queue.Push(Node{name: root, weight: 0})

	for {
		actualNode := queue.Pop()
		if actualNode == nil {
			break
		}

		if distances[actualNode.(Node).name] < actualNode.(Node).weight {
			continue
		}

		for _, node := range graph[actualNode.(Node).name] {

			currentWeight := actualNode.(Node).weight + node.weight
			distancesWeight := distances[node.name]

			if currentWeight < distancesWeight {
				distances[node.name] = currentWeight
				queue.Push(Node{name: node.name, weight: currentWeight})
			}

		}
	}

	return distances
}
