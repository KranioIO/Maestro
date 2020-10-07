package parser

import (
	"app/common"
	"app/orchestration/symphony"
	"errors"
)

// GenerateDAGfromDict creates a directed acyclic graph object from the dict generated from yml file
func GenerateDAGfromDict(dict common.Dict) (*symphony.DAG, error) {
	if err := symphonyDictIsNotValid(dict); err != nil {
		return nil, err
	}

	nodes, err := generateDagNodesFromTasks(dict["tasks"].(common.Dict))

	if err != nil {
		return nil, err
	}

	dag := symphony.DAG{
		Name:  dict["symphony"].(string),
		Nodes: nodes,
	}

	if err := dag.GenerateExecutionTree(); err != nil { // TODO: take a better look in memory leak
		return nil, err
	}

	return &dag, nil
}

/* TODO:
- improve error detection passing the name of wrong field
- validate field type as well
*/
func symphonyDictIsNotValid(dict common.Dict) error {
	var validations [3]bool

	_, validations[0] = dict["symphony"]
	_, validations[1] = dict["trigger"]
	_, validations[2] = dict["tasks"]

	for _, exist := range validations {
		if !exist {
			return errors.New("yml schema is not correct")
		}
	}

	return nil
}

// TODO: implement task validation
func tasksDictIsNotValid(dict common.Dict) error {
	return nil
}

func generateDagNodesFromTasks(tasks common.Dict) ([]*symphony.Node, error) {
	var nodes []*symphony.Node

	if err := tasksDictIsNotValid(tasks); err != nil {
		return nil, err
	}

	for name, definition := range tasks {
		node := &symphony.Node{Name: name.(string)}
		nodes = append(nodes, node)

		taskDict := definition.(common.Dict)
		node.Type = taskDict["type"].(string)

		if payload, exist := taskDict["payload"]; exist {
			node.Payload = payload.(common.Dict)
		}

		if dependency, exists := taskDict["depends"]; exists {
			switch dependency.(type) {
			case string:
				node.Depends = []string{dependency.(string)}

			case []interface{}:
				dependencies := dependency.([]interface{})
				node.Depends = make([]string, 0, len(dependencies))

				for _, taskName := range dependencies {
					node.Depends = append(node.Depends, taskName.(string))
				}
			}
		}
	}

	return nodes, nil
}
