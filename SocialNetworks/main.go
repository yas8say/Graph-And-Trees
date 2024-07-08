package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Graph structure using adjacency list representation
type Graph struct {
	nodes map[string][]string
}

// NewGraph initializes a new Graph
func NewGraph() *Graph {
	return &Graph{nodes: make(map[string][]string)}
}

// AddEdge adds an edge between two nodes (from and to) in the graph
func (g *Graph) AddEdge(from, to string) {
	g.nodes[from] = append(g.nodes[from], to) // Add 'to' to the adjacency list of 'from'
	g.nodes[to] = append(g.nodes[to], from)   // Add 'from' to the adjacency list of 'to' (for undirected graph)
}

// PrintGraph prints the adjacency list of the graph
func (g *Graph) PrintGraph() {
	for node, edges := range g.nodes {
		fmt.Printf("%s: %v\n", node, edges) // Print each node and its list of adjacent nodes
	}
}

func main() {
	graph := NewGraph()                 // Create a new graph
	reader := bufio.NewReader(os.Stdin) // Create a reader to read input from the user

	// Loop to take user input until 'done' is entered
	for {
		fmt.Print("Enter an edge (from to) or type 'done' to finish: ")
		input, err := reader.ReadString('\n') // Read input from the user
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}
		input = strings.TrimSpace(input) // Trim any extra whitespace
		if input == "done" {
			break // Exit the loop if the user types 'done'
		}
		parts := strings.Split(input, " ") // Split the input into 'from' and 'to'
		if len(parts) != 2 {
			fmt.Println("Invalid input. Please enter 'from' and 'to' separated by a space.")
			continue // Continue the loop for valid inputs
		}
		from, to := parts[0], parts[1]
		fmt.Printf("Adding edge from %s to %s\n", from, to) // Debug: print the edge being added
		graph.AddEdge(from, to)                             // Add the edge to the graph
	}

	fmt.Println("Final graph:")
	graph.PrintGraph() // Print the final adjacency list of the graph
}
