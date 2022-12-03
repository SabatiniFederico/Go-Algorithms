package prim

import (
	"fmt"
	"testing"
)

func TestPrim(t *testing.T) {

	graph := make(map[string][]Edge)

	graph["saba"] = []Edge{{from: "saba", to: "mega", weight: 100}, {from: "saba", to: "tiki", weight: 800}}
	graph["mega"] = []Edge{{from: "mega", to: "saba", weight: 100}, {from: "mega", to: "tiki", weight: 200}}
	graph["tiki"] = []Edge{{from: "tiki", to: "saba", weight: 800}, {from: "tiki", to: "mega", weight: 200}}

	agm := prim(graph, "saba")
	fmt.Println(agm)
}

func TestPrimComplicated(t *testing.T) {
	graph := make(map[string][]Edge)

	graph["a"] = []Edge{
		{from: "a", to: "b", weight: 5},
		{from: "a", to: "c", weight: 5},
		{from: "a", to: "d", weight: 5}}
	graph["b"] = []Edge{
		{from: "b", to: "a", weight: 5},
		{from: "b", to: "d", weight: 1}}
	graph["c"] = []Edge{
		{from: "c", to: "a", weight: 5},
		{from: "c", to: "d", weight: 1}}
	graph["d"] = []Edge{
		{from: "d", to: "a", weight: 5},
		{from: "d", to: "b", weight: 1},
		{from: "d", to: "c", weight: 1}}

	agm := prim(graph, "a")
	fmt.Println(agm)
}

//Prim
//map[a:[{a b 5}] b:[{a b 5} {b d 1}] c:[{d c 1}] d:[{b d 1} {d c 1}]]

//Dijkstra
//map[a:0 b:5 c:5 d:5]
//map[a:5 b:0 c:2 d:1]
