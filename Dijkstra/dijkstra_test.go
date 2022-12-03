package dijkstra

import (
	"fmt"
	"testing"

	"github.com/x1m3/priorityQueue"
)

func TestDijkstra(t *testing.T) {

	graph := make(map[string][]Node)

	graph["saba"] = []Node{{name: "mega", weight: 100}, {name: "naval", weight: 500}}
	graph["mega"] = []Node{{name: "tiki", weight: 5}, {name: "facu", weight: 1000}}
	graph["tiki"] = []Node{}
	graph["facu"] = []Node{{name: "saba", weight: 200}}
	graph["nava"] = []Node{{name: "facu", weight: 300}}
	graph["inaccesible"] = []Node{}

	root := "saba"
	distances := dijkstra(graph, root)
	println(distances)

}

func TestPriorityQueueSimple(t *testing.T) {

	list := priorityQueue.New()
	list.Push(Node{name: "saba", weight: 200})
	list.Push(Node{name: "mega", weight: 500})
	list.Push(Node{name: "facu", weight: 300})

	for {
		r := list.Pop()
		if r == nil {
			break
		}
		fmt.Println(r)
	}
}
