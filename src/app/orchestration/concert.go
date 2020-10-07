package orchestration

import (
	"app/orchestration/parser"
	"app/orchestration/symphony"
	"app/orchestration/trigger"
	"app/storage"
	"log"
)

// Dags is the reference to all loaded symphonies running on the maestro
var Dags []*symphony.DAG

// Triggers ..
var Triggers map[string]*trigger.Trigger

// BeginConcert reproduce the initial setup steps of the application
func BeginConcert() chan *symphony.DagExecution {
	filePaths := storage.GetAllSymphoniesFilenamesInFilesystem()
	Triggers = make(map[string]*trigger.Trigger)
	Dags = make([]*symphony.DAG, 0, len(filePaths))
	executionLogger := trigger.GetLogger()

	for _, filePath := range filePaths {
		dict := storage.ReadYmlFile(filePath)

		var dag *symphony.DAG
		var err error

		if dag, err = parser.GenerateDAGfromDict(dict); err != nil {
			log.Println(err) // send error to a channel
			continue
		}

		if valid, err := dag.IsValid(); err != nil || !valid {
			log.Println(err) // send error to a channel
			continue
		}

		dagTrigger, err := parser.GenerateTriggerFromDict(dict)

		if err != nil {
			log.Println(err) // send error to a channel
			continue
		}

		Triggers[dag.Name] = dagTrigger
		dagTrigger.Assign(dag)

		Dags = append(Dags, dag)
	}

	return executionLogger
}

// receive error stream as parameter to send validation errors to the webserver page
