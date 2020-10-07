package symphony

import (
	"fmt"
)

// DAG (directed acyclic graph) is the representation of mathematical graph model of a flow
type DAG struct {
	Active bool
	Name   string
	Nodes  []*Node
	Root   []*Node
}

// Node is a part of a DAG, have his own identity and a link for the next Node to be executed
type Node struct {
	Name    string
	Type    string
	Next    []*Node
	Prev    []*Node
	Depends []string
	Payload map[interface{}]interface{}
}

// GenerateExecutionTree update the next and previus nodes reference on each node
func (dag *DAG) GenerateExecutionTree() error {
	var nodeIndexer = make(map[string]*Node)

	for _, node := range dag.Nodes {
		nodeIndexer[node.Name] = node
	}

	for _, node := range dag.Nodes {
		if dependencies := node.Depends; len(dependencies) > 0 {
			for _, dependency := range dependencies {
				if _, exist := nodeIndexer[dependency]; !exist {
					return fmt.Errorf("node %s depends on a non-existent node %s", node.Name, dependency)
				}

				nodeIndexer[dependency].Next = append(nodeIndexer[dependency].Next, node)
				node.Prev = append(node.Prev, nodeIndexer[dependency])
			}
		} else {
			dag.Root = append(dag.Root, node)
		}
	}

	return nil
}

// IsValid execute a series of validations over the dag to prevents execution problems or deadlocks
func (dag *DAG) IsValid() (bool, error) {
	var visited = make(map[string]bool)

	for _, node := range dag.Nodes {
		visited[node.Name] = false
	}

	if len(dag.Root) == 0 || executionLoopFound(dag.Root, visited) {
		return false, fmt.Errorf("loop found on the dag %s by a cross dependency cycle", dag.Name)
	}

	return true, nil
}

// depth first search algorithm for cycle identification
func executionLoopFound(next []*Node, visited map[string]bool) bool {
	var nextNodes = make(map[string]*Node)

	for _, node := range next {
		if visited[node.Name] {
			return true
		}

		visited[node.Name] = true

		for _, queuedNode := range node.Next {
			nextNodes[node.Name] = queuedNode
		}
	}

	if len(nextNodes) == 0 {
		return false
	}

	executionQueue := make([]*Node, len(nextNodes))
	idx := 0

	for _, node := range nextNodes {
		executionQueue[idx] = node
		idx++
	}

	return executionLoopFound(executionQueue, visited)
}
