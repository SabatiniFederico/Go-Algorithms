package prim

import "github.com/x1m3/priorityQueue"

type Edge struct {
	from   string
	to     string
	weight int
}

// I need to define what priority means in a struct
func (i Edge) HigherPriorityThan(other priorityQueue.Interface) bool {
	return i.weight < other.(Edge).weight
}

//root should exist in graph
func prim(graph map[string][]Edge, root string) map[string][]Edge {

	agm := make(map[string][]Edge)
	queue := priorityQueue.New()

	for key := range graph {
		agm[key] = []Edge{}
	}

	for _, edge := range graph[root] {
		queue.Push(edge)
	}

	for {
		poppedEdge := queue.Pop()
		if poppedEdge == nil {
			break
		}
		minEdge := poppedEdge.(Edge)

		if len(agm[minEdge.to]) == 0 {
			agm[minEdge.from] = append(agm[minEdge.from], minEdge)
			agm[minEdge.to] = append(agm[minEdge.to], minEdge)
		}

		//AGM tiene actualNode -> node con un peso
		for _, node := range graph[minEdge.to] {

			if len(agm[node.to]) == 0 {
				queue.Push(node)
			}
		}
	}
	return agm
}
