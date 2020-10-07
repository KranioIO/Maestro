package storage

import (
	"app/orchestration/parser"
	"app/orchestration/symphony"
	"log"
)

// LogExecutions will listen the execution channel and save everything on mongodb
func LogExecutions(executionChanel chan *symphony.DagExecution) {
	for dagExecution := range executionChanel {
		execution := parser.ExecutionToDict(dagExecution)

		if err := InsertData("executions", execution); err != nil {
			log.Println(execution)
			log.Println("insert on db failed")
		}
	}
}
