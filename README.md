# dagx

A lightweight, zero-dependency Directed Acyclic Graph (DAG) library for Go.

## Features

- Simple and intuitive API
- Zero external dependencies
- Cycle detection
- Topological sorting (Kahn's algorithm)
- Parent/child relationship traversal
- In-degree calculation
- JSON-serializable Node and Edge types

## Installation

```bash
go get github.com/oshturhq/dagx
```

## Usage

### Creating a DAG

```go
package main

import (
    "fmt"
    "github.com/oshturhq/dagx"
)

func main() {
    dag := dagx.NewDAG()

    // Add nodes
    dag.AddNode("A")
    dag.AddNode("B")
    dag.AddNode("C")

    // Add edges (A -> B -> C)
    dag.AddEdge("A", "B")
    dag.AddEdge("B", "C")
}
```

### Topological Sorting

```go
order, err := dag.TopologicalOrder()
if err != nil {
    if err == dagx.ErrCycleDetected {
        fmt.Println("Graph contains a cycle!")
    }
    if err == dagx.ErrEmptyGraph {
        fmt.Println("Graph is empty!")
    }
}
fmt.Println(order) // Output: [A B C]
```

### Traversing Relationships

```go
// Get all children of a node
children := dag.Children("A") // Returns: ["B"]

// Get all parents of a node
parents := dag.Parents("C") // Returns: ["B"]

// Get in-degree (number of incoming edges)
inDegree := dag.InDegree("B") // Returns: 1
```

### Cycle Detection

```go
dag := dagx.NewDAG()
dag.AddEdge("A", "B")
dag.AddEdge("B", "C")
dag.AddEdge("C", "A") // Creates a cycle

if dag.HasCycle() {
    fmt.Println("Cycle detected!")
}
```

### Listing Nodes and Edges

```go
// Get all nodes
nodes := dag.Nodes() // Returns: ["A", "B", "C"]

// Get all edges
edges := dag.Edges() // Returns: [{From: "A", To: "B"}, {From: "B", To: "C"}]
```

## API Reference

### Types

- `DAG` - The main directed acyclic graph structure
- `Node` - Represents a graph node with an ID field
- `Edge` - Represents a directed edge with From and To fields

### Methods

| Method | Description |
|--------|-------------|
| `NewDAG()` | Creates a new empty DAG |
| `AddNode(id string)` | Adds a node to the graph |
| `AddEdge(from, to string)` | Adds a directed edge between two nodes |
| `Nodes() []string` | Returns all node IDs |
| `Edges() []Edge` | Returns all edges |
| `Parents(id string) []string` | Returns parent node IDs |
| `Children(id string) []string` | Returns child node IDs |
| `InDegree(id string) int` | Returns the number of incoming edges |
| `TopologicalOrder() ([]string, error)` | Returns nodes in topological order |
| `HasCycle() bool` | Returns true if the graph contains a cycle |

### Errors

- `ErrCycleDetected` - Returned when a cycle is detected in the graph
- `ErrEmptyGraph` - Returned when attempting operations on an empty graph

## Use Cases

- Task scheduling and dependency resolution
- Build systems and compilation order
- Data pipeline orchestration
- Workflow engines
- Package dependency management
