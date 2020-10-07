package parser_test

import (
	"app/common"
	"app/orchestration/parser"
	"testing"
)

func TestGenerateDAGfromDict(t *testing.T) {
	dagDict := common.Dict{
		"symphony": "test-dict-parser",
		"schedule": "* * * * *",
		"tasks": common.Dict{
			"task-A": common.Dict{},
			"task-B": common.Dict{
				"depends": "task-A", // <- string
			},
			"task-C": common.Dict{
				"depends": []interface{}{"task-A"}, // <- single []interface{}
			},
			"task-D": common.Dict{
				"depends": []interface{}{"task-B", "task-C"}, // <- multiple []interface{}
			},
		},
	}

	wrongDagDict := common.Dict{
		"symphony": "test-dict-parser",
		"schedule": "* * * * *",
	}

	if _, err := parser.GenerateDAGfromDict(dagDict); err != nil {
		t.Error(err)
	}

	if _, err := parser.GenerateDAGfromDict(wrongDagDict); err == nil {
		t.Error(err)
	}
}
