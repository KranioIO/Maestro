package parser

import "app/orchestration/symphony"

// NodesToMap return Nodes as a list of maps
func NodesToMap(dag *symphony.DAG) []map[string]interface{} {
	nodeMaps := make([]map[string]interface{}, len(dag.Nodes))

	for idx, node := range dag.Nodes {
		nodeMap := make(map[string]interface{})

		nodeMap["name"] = node.Name
		nodeMap["depends"] = node.Depends

		nodeMaps[idx] = nodeMap
	}

	return nodeMaps
}

// func (instance *DagInstance) parseToDbRegistry() storage.DagExecutionRegistry {
// 	var nodes []storage.NodeExecutionRegistry

// 	for _, ni := range instance.allNodes {
// 		nodes = append(nodes, storage.NodeExecutionRegistry{
// 			NodeID:  ni.node.name,
// 			Started: ni.startedAt,
// 			Ended:   ni.endedAt,
// 		})
// 	}

// 	return storage.DagExecutionRegistry{
// 		ExecutionTime: instance.executionTime,
// 		DagID:         instance.dag.name,
// 		Nodes:         nodes,
// 	}
// }
