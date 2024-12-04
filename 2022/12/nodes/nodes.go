package nodes

import (
	"log"
	"strings"
)

type Node struct {
	UpNeighbour   *Node
	DownNeighbour *Node
	RighNeighbour *Node
	LeftNeighbour *Node
	Height        string
}

func PrintNodes(nodes [][]*Node) (output string) {
	for rowIndex, nodesLine := range nodes {
		verticalConnections := ""
		for columnIndex, node := range nodesLine {
			horisontalConnection := " "
			if rowIndex < len(nodes)-1 && node.DownNeighbour != nil {
				verticalConnections += "| "
			} else {
				verticalConnections += "  "
			}
			if columnIndex < len(nodesLine)-1 && node.RighNeighbour != nil {
				horisontalConnection = "-"
			}
			output += node.Height + horisontalConnection
		}
		output += "\n" + verticalConnections + "\n"
	}
	return output
}

func (node *Node) GetHeight() int {
	alphabet := "SabcdefghijklmnopqrstuvwxyzE"
	return strings.Index(alphabet, node.Height)
}
func GetNodeHeight(letter string) int {
	alphabet := "SabcdefghijklmnopqrstuvwxyzE"
	return strings.Index(alphabet, letter)
}
func AreNodesConnected(node1, node2 *Node) bool {
	if node1.GetHeight()+1 == node2.GetHeight() || node1.GetHeight()-1 == node2.GetHeight() || node1.Height == node2.Height {
		return true
	}
	return false
}

func ConnectNodes(nodes [][]*Node) Graph {
	connectionGraph := newGraph()
	for rowIndex, nodesLine := range nodes {
		for columnIndex, node := range nodesLine {

			if columnIndex <= len(nodesLine)-2 && AreNodesConnected(node, nodesLine[columnIndex+1]) {
				node.RighNeighbour = nodesLine[columnIndex+1]
				nodesLine[columnIndex+1].LeftNeighbour = node
				connectionGraph = connectionGraph.link(node, nodesLine[columnIndex+1])
				if node.Height == "S" {
					log.Printf("graph for S: %+v\n", connectionGraph[node])
				}
			}
			if rowIndex <= len(nodes)-2 && AreNodesConnected(node, nodes[rowIndex+1][columnIndex]) {
				node.DownNeighbour = nodes[rowIndex+1][columnIndex]
				nodes[rowIndex+1][columnIndex].UpNeighbour = node
				connectionGraph = connectionGraph.link(node, nodes[rowIndex+1][columnIndex])
			}
		}
	}
	return connectionGraph
}

func FindByLetter(nodes [][]*Node, name string) *Node {
	for _, nodesLine := range nodes {
		for _, node := range nodesLine {
			if node.Height == name {
				return node
			}
		}
	}
	return nil
}

type Graph map[*Node][]*Node

func newGraph() Graph {
	return make(map[*Node][]*Node)
}

func FindNodeInList(list []*Node, findNode *Node) bool {
	for _, node := range list {
		if node == findNode {
			return true
		}
	}
	return false
}

// link creates a bi-directed edge between nodes a and b.
func (g Graph) link(a, b *Node) Graph {
	g[a] = append(g[a], b)
	g[b] = append(g[b], a)
	if a.Height == "S" {
		log.Printf("in link graph for S: %+v", g[a])
	}
	return g
}

// Neighbours returns the neighbour nodes of node n in the Graph.
func (g Graph) Neighbours(n *Node) []*Node {
	nodes := []*Node{}
	for _, node := range []*Node{n.UpNeighbour, n.LeftNeighbour, n.RighNeighbour, n.LeftNeighbour} {
		if node != nil {
			nodes = append(nodes, node)
		}
	}
	return nodes
}

// nodeDist is our cost function. We use points as nodes, so we
// calculate their Euclidean distance.
func NodeDist(p, q *Node) int {
	if p.Height == "S" || p.Height == "E" || q.Height == "S" || q.Height == "E" {
		return 0
	}
	return p.GetHeight() - q.GetHeight()
}
