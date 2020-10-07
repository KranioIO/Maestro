package symphony_test

import (
	"app/orchestration/symphony"
	"log"
	"testing"
)

func TestDagExecutionTreeGeneration(t *testing.T) {
	// No wrong dependencies
	correctDAG := symphony.DAG{Nodes: []*symphony.Node{
		&symphony.Node{Name: "A"},
		&symphony.Node{Name: "B", Depends: []string{"A"}},
		&symphony.Node{Name: "C", Depends: []string{"B"}},
	}}

	// C dependes on inexistent D Node
	incorrectDAG := symphony.DAG{Nodes: []*symphony.Node{
		&symphony.Node{Name: "A"},
		&symphony.Node{Name: "B", Depends: []string{"A"}},
		&symphony.Node{Name: "C", Depends: []string{"D"}},
	}}

	if err := correctDAG.GenerateExecutionTree(); err != nil {
		t.Error("DAG should generate execution Tree without errors")
	}

	if err := incorrectDAG.GenerateExecutionTree(); err == nil {
		log.Println(err)
		t.Error("DAG shouldn't generate execution")
	}
}

func TestDagValidation(t *testing.T) {
	// Correct execution no cycles or cross dependencies
	// A -> B -> C
	dagCorrectLineModel := symphony.DAG{Nodes: []*symphony.Node{
		&symphony.Node{Name: "A"},
		&symphony.Node{Name: "B", Depends: []string{"A"}},
		&symphony.Node{Name: "C", Depends: []string{"B"}},
	}}

	// Correct execution but E is call more than once (only execute when everyone before is complete, no cycles or cross dependencies
	//    /-> B -> E
	// A ---> C -> E
	//    \-> D -> E
	dagCorrectTreeModel := symphony.DAG{Nodes: []*symphony.Node{
		&symphony.Node{Name: "A"},
		&symphony.Node{Name: "B", Depends: []string{"A"}},
		&symphony.Node{Name: "C", Depends: []string{"A"}},
		&symphony.Node{Name: "D", Depends: []string{"A"}},
		&symphony.Node{Name: "E", Depends: []string{"B", "C", "D"}},
	}}

	// A cycle, didn't generate a root node, is never executed
	// A -> B -> C -> A -> B -> C -> A ...
	dagWithCycle := symphony.DAG{Nodes: []*symphony.Node{
		&symphony.Node{Name: "A", Depends: []string{"C"}},
		&symphony.Node{Name: "B", Depends: []string{"A"}},
		&symphony.Node{Name: "C", Depends: []string{"B"}},
	}}

	// The execution gets stuck cause the dependency will never be executed (indirect cross dependency) (deadlock)
	// A -> B -> C -> D
	//      ^---------^
	dagUnreachableDependency := symphony.DAG{Nodes: []*symphony.Node{
		&symphony.Node{Name: "A"},
		&symphony.Node{Name: "B", Depends: []string{"A", "D"}},
		&symphony.Node{Name: "C", Depends: []string{"B"}},
		&symphony.Node{Name: "D", Depends: []string{"C"}},
	}}

	// the execution gets stuck cause will never execute the dependencies (direct cross dependency)
	// A -> B <-> C
	dagWithCrossDependency := symphony.DAG{Nodes: []*symphony.Node{
		&symphony.Node{Name: "A"},
		&symphony.Node{Name: "B", Depends: []string{"A", "C"}},
		&symphony.Node{Name: "C", Depends: []string{"B"}},
	}}

	dagUnreachableDependency.GenerateExecutionTree()
	dagWithCrossDependency.GenerateExecutionTree()
	dagCorrectTreeModel.GenerateExecutionTree()
	dagCorrectLineModel.GenerateExecutionTree()
	dagWithCycle.GenerateExecutionTree()

	if valid, _ := dagCorrectLineModel.IsValid(); !valid {
		t.Error("DAG line model should be valid")
	}

	if valid, _ := dagCorrectTreeModel.IsValid(); !valid {
		t.Error("DAG tree model should be valid")
	}

	if valid, _ := dagWithCycle.IsValid(); valid {
		t.Error("DAG with cycle shouldn't be valid")
	}

	if valid, _ := dagWithCrossDependency.IsValid(); valid {
		t.Error("DAG with direct cross dependency shouldn't be valid")
	}

	if valid, _ := dagUnreachableDependency.IsValid(); valid {
		t.Error("DAG with indirect cross dependency shouldn't be valid")
	}
}

func TestDagExecution(t *testing.T) {
	dag := symphony.DAG{Nodes: []*symphony.Node{
		&symphony.Node{Name: "A", Type: "dummy"},
		&symphony.Node{Name: "B", Type: "dummy", Depends: []string{"A"}},
		&symphony.Node{Name: "C", Type: "dummy", Depends: []string{"A"}},
		&symphony.Node{Name: "D", Type: "dummy", Depends: []string{"A"}},
		&symphony.Node{Name: "E", Type: "dummy", Depends: []string{"B", "C", "D"}},
	}}

	dag.GenerateExecutionTree()
	dagExecution := symphony.CreateDagExecution(&dag)
	dagExecution.Run()

	for _, node := range dagExecution.Nodes {
		if logs := node.Task.GetLogs(); len(logs) > 1 {
			t.Error("Node executed more then once")
		}

		// log.Println(node.Node.Name, node.Task.GetLogs())
	}
}
