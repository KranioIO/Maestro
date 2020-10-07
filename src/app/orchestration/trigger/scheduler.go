package trigger

import (
	"sync"

	"github.com/robfig/cron/v3"
)

// Scheduler is a brigde for dag and cron management
type Scheduler struct {
	cronManager *cron.Cron
}

var once sync.Once
var scheduler *Scheduler

// GetScheduler works with singleton approach and will return an scheduler reference
func GetScheduler() *Scheduler {
	once.Do(func() {
		scheduler = &Scheduler{cronManager: cron.New()}
		scheduler.cronManager.Start()
	})

	return scheduler
}

func (scheduler *Scheduler) addTrigger(trigger *Trigger) {
	entryID, _ := scheduler.cronManager.AddFunc(trigger.Info, trigger.Fire) // everyone has it's own goroutine when called
	trigger.Cron = entryID
}

// // GetExecutionTime ..
// func GetExecutionTime(dag *symphony.DAG) time.Time {
// 	return scheduler.cronManager.Entry(dag.entryID).Prev
// }

// func getNextExecution(dag *symphony.DAG) time.Time {
// 	return scheduler.cronManager.Entry(dag.entryID).Next
// }
