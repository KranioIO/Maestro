package parser

import (
	"app/orchestration/symphony"
)

// ExecutionToDict will extract important info from execution and pass to a map for external consumption
func ExecutionToDict(dagExec *symphony.DagExecution) map[string]interface{} {
	dagExecDict := make(map[string]interface{})
	nodes := make([]map[string]interface{}, 0, len(dagExec.Nodes))

	for _, node := range dagExec.Nodes {
		nodes = append(nodes, map[string]interface{}{
			"ExecutionID": node.ExecutionID,
			"Name":        node.Node.Name,
			"Type":        node.Node.Type,
			"Logs":        node.Task.GetLogs(),
			"State":       node.State,
		})
	}

	dagExecDict["ExecutionID"] = dagExec.ExecutionID
	dagExecDict["Name"] = dagExec.Dag.Name
	dagExecDict["Nodes"] = nodes

	return dagExecDict
}
