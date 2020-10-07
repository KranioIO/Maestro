package symphony

import (
	"app/orchestration/operator"
	"log"
	"sync"
	"time"

	"github.com/google/uuid"
)

const (
	nodeQueued  = 0
	nodeRunning = iota
	nodeSuccess = iota
	nodeWarning = iota
	nodeSkipped = iota
	nodeErr     = iota
)

// DagExecution is an execution instance of a DAG
type DagExecution struct {
	ExecutionID   uuid.UUID
	ExecutionTime time.Time
	Nodes         []*NodeExecution
	Root          []*NodeExecution
	Dag           *DAG
}

// NodeExecution is an execution instance of a Node
type NodeExecution struct {
	ExecutionID uuid.UUID
	Task        operator.Operator
	Prev        []*NodeExecution
	Next        []*NodeExecution
	Node        *Node
	State       int
}

// CreateDagExecution parse the DAG to an execution instance with temporal information
func CreateDagExecution(dag *DAG) *DagExecution {
	nodeExecIndexer := make(map[string]*NodeExecution)
	dagExec := DagExecution{
		Dag:         dag,
		Root:        make([]*NodeExecution, 0, len(dag.Root)), // prevents append re-allocation
		Nodes:       make([]*NodeExecution, 0, len(dag.Nodes)),
		ExecutionID: uuid.New(),
	}

	for _, node := range dag.Nodes {
		nodeExecIndexer[node.Name] = parseNode(node)
		dagExec.Nodes = append(dagExec.Nodes, nodeExecIndexer[node.Name])
	}

	for _, node := range dag.Root {
		dagExec.Root = append(dagExec.Root, nodeExecIndexer[node.Name])
	}

	for _, node := range dag.Nodes {
		nodeExecIndexer[node.Name].Next = make([]*NodeExecution, 0, len(node.Next))

		for _, next := range node.Next {
			nodeExecIndexer[node.Name].Next = append(nodeExecIndexer[node.Name].Next, nodeExecIndexer[next.Name])
			nodeExecIndexer[next.Name].Prev = append(nodeExecIndexer[next.Name].Prev, nodeExecIndexer[node.Name])
		}
	}

	return &dagExec
}

type operatorCreationFunc func(map[interface{}]interface{}) operator.Operator

func parseNode(node *Node) *NodeExecution {
	taskCreator := map[string]operatorCreationFunc{
		"dummy":  operator.CreateDummy,
		"lambda": operator.CreateLambda,
	}

	return &NodeExecution{
		ExecutionID: uuid.New(),
		Node:        node,
		Task:        taskCreator[node.Type](node.Payload),
	}
}

// Run function starts the execution tree for an specific Symphony
func (dagExec *DagExecution) Run() {
	log.Printf("Starting %s pipeline - execution id: %s", dagExec.Dag.Name, dagExec.ExecutionID) // TODO add to a dag log

	for next := runNodes(dagExec.Root); len(next) > 0; next = runNodes(next) {
	}
}

func runNodes(nodes []*NodeExecution) []*NodeExecution {
	var next = make(map[string]*NodeExecution)
	var wg sync.WaitGroup
	var mu sync.Mutex

	for _, node := range nodes {
		if node.previousCompleted() {
			wg.Add(1)

			go func(nodeExec *NodeExecution) {
				defer wg.Done()

				if err := nodeExec.Task.Execute(); err == nil {
					nodeExec.State = nodeSuccess

					for _, downstream := range nodeExec.Next {
						mu.Lock()
						next[downstream.Node.Name] = downstream
						mu.Unlock()
					}
				} else {
					// error treatment and logs
					setDownstreamStatus(nodeExec.Next, nodeSkipped)
				}
			}(node)
		}
	}

	wg.Wait()
	nextNodes := make([]*NodeExecution, 0, len(next))

	for _, nodeExec := range next {
		nextNodes = append(nextNodes, nodeExec)
	}

	return nextNodes
}

func (nodeExec *NodeExecution) previousCompleted() bool {
	upstreamComplete := true

	for _, previous := range nodeExec.Prev {
		upstreamComplete = upstreamComplete && (previous.State == nodeSuccess || previous.State == nodeErr)
	}

	return upstreamComplete
}

func setDownstreamStatus(nodes []*NodeExecution, status int) {
	// code..
}
