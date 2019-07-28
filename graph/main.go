package main

import (
	"fmt"

	"github.com/zoumo/chaos/graph/set"
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/simple"
	"gonum.org/v1/gonum/graph/topo"
)

type Directed struct {
	Graph       graph.Directed
	StartPoints []graph.Node
}

type Node struct {
	id   int64
	name string
}

func (n Node) ID() int64 {
	return n.id
}

func main() {
	g := simple.NewDirectedGraph()
	// es := []simple.Edge{
	// 	{F: Node{1, ""}, T: Node{2, ""}},
	// 	{F: Node{1, ""}, T: Node{3, ""}},
	// 	{F: Node{2, ""}, T: Node{4, ""}},
	// 	{F: Node{3, ""}, T: Node{4, ""}},
	// 	{F: Node{4, ""}, T: Node{5, ""}},
	// 	{F: Node{5, ""}, T: Node{2, ""}},
	// 	{F: Node{6, ""}, T: Node{7, ""}},
	// 	{F: Node{7, ""}, T: Node{6, ""}},
	// 	{F: Node{8, ""}, T: Node{4, ""}},
	// }
	es := []simple.Edge{
		{F: Node{1, ""}, T: Node{2, ""}},
		{F: Node{2, ""}, T: Node{3, ""}},
		{F: Node{3, ""}, T: Node{4, ""}},
		{F: Node{3, ""}, T: Node{5, ""}},
		{F: Node{4, ""}, T: Node{5, ""}},
		{F: Node{5, ""}, T: Node{1, ""}},
	}
	g.AddNode(Node{9, ""})
	g.AddNode(Node{10, ""})
	g.AddNode(Node{11, ""})

	for _, e := range es {
		// fmt.Println(e.To(), e.From())
		// ne := g.NewEdge(e.To(), e.From())
		ne := g.NewEdge(e.From(), e.To())
		g.SetEdge(ne)
	}
	g, startedNodes, _ := separateIsolatedNodes(g)
	fmt.Printf("started Nodes: %v\n", startedNodes)
	gcopy := simple.NewDirectedGraph()
	graph.Copy(gcopy, g)
	// kahn_cycles(gcopy, startedNodes)

	ret := topo.TarjanSCC(gcopy)
	for i, ns := range ret {
		// for j, n := range ns {
		// }
		// if len(ns) == 1 {
		// 	continue
		// }
		fmt.Println(i, ns)
	}
}

// kahn 算法
func kahn_cycles(g *simple.DirectedGraph, started set.Nodes) {

	if len(started) == 0 {
		return
	}

	for {
		next := set.Nodes{}
		entryPoint := set.Nodes{}
		// started 中存放的是入度为 0 的点，找到这个点的所有后驱
		// 如果后驱的入度为 1, 则说明这点被删除后，后驱的入度会变成 0
		for _, n := range started {
			for _, nn := range graph.NodesOf(g.From(n.ID())) {
				if g.To(nn.ID()).Len() == 1 {
					//
					next.Add(nn)
				} else {
					entryPoint.Add(nn)
				}
			}
			g.RemoveNode(n.ID())
		}

		if len(next) == 0 {
			started = entryPoint
			break
		}

		started = next
	}

	leftNodes := g.Nodes()
	if leftNodes.Len() == 0 {
		return
	}

	fmt.Println(leftNodes)
	fmt.Println(started)
}

func dfs(g graph.Directed, node graph.Node) {

}

func separateIsolatedNodes(g *simple.DirectedGraph) (*simple.DirectedGraph, set.Nodes, set.Nodes) {
	startNodes := set.Nodes{}
	isolatedNodes := set.Nodes{}
	for _, n := range graph.NodesOf(g.Nodes()) {
		if g.To(n.ID()).Len() == 0 {
			if g.From(n.ID()).Len() == 0 {
				isolatedNodes.Add(n)
				g.RemoveNode(n.ID())
				continue
			}
			startNodes.Add(n)
		}
	}

	return g, startNodes, isolatedNodes

}

// func splitGraph(g graph.Directed) {

// 	visit := func(g graph.Directed, visited goset.Set, nodes goset.Set) goset.Set {
// 		next := goset.NewSet()
// 		nodes.Range(func(i int, e interface{}) bool {
// 			node := e.(graph.Node)
// 			for _, n := range graph.NodesOf(g.From(node.ID())) {
// 				if visited.Contains(n) {
// 					continue
// 				}
// 				next.Add(n)
// 			}
// 			return true
// 		})
// 		return next
// 	}

// 	dfs := func(g graph.Directed, node graph.Node) goset.Set {
// 		visited := goset.NewSet(node)
// 		next := visit(g, visited, goset.NewSet(node))
// 		visited.Extend(next)
// 		for next.Len() > 0 {
// 			next = visit(g, visited, next)
// 			visited.Extend(next)
// 		}

// 		return visited
// 	}

// 	startNodes := []graph.Node{}
// 	for _, n := range graph.NodesOf(g.Nodes()) {
// 		if g.To(n.ID()).Len() == 0 {
// 			startNodes = append(startNodes, n)
// 			set := dfs(g, n)
// 			fmt.Println(set.Elements())
// 		}
// 	}
// }
