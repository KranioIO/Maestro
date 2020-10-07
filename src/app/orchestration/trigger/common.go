package trigger

import (
	"app/orchestration/symphony"
	"sync"
)

var logger chan *symphony.DagExecution
var onceLogger sync.Once

// Trigger contains information of how an dag should be triggered
type Trigger struct {
	Type string
	Info string
	Dag  *symphony.DAG
	Cron interface{}
}

// GetLogger works with singleton approach and will return an scheduler reference
func GetLogger() chan *symphony.DagExecution {
	onceLogger.Do(func() {
		logger = make(chan *symphony.DagExecution)
	})

	return logger
}

// Assign will add the dag to be called from the respective trigger type
func (trigger *Trigger) Assign(dag *symphony.DAG) {
	trigger.Dag = dag

	if trigger.Type == "scheduler" {
		GetScheduler().addTrigger(trigger)
	}
}

// Fire will start an execution for the referenced trigger
func (trigger *Trigger) Fire() {
	dagExecution := symphony.CreateDagExecution(trigger.Dag)
	dagExecution.Run()

	logger <- dagExecution
}
