package operator

import (
	"fmt"
	"math/rand"
	"time"
)

// DummyOperator used just for symphony development
type DummyOperator struct {
	StartedAt time.Time
	EndedAt   time.Time
	Logs      []string
}

// CreateDummy returns a dummy operator instance
func CreateDummy(payload map[interface{}]interface{}) Operator {
	return &DummyOperator{}
}

// Execute will perform the goal of the operator
func (dummyOp *DummyOperator) Execute() error {
	dummyOp.StartedAt = time.Now()
	randDuration := rand.Intn(10)

	time.Sleep(time.Duration(randDuration) * time.Second)

	dummyOp.Logs = append(dummyOp.Logs, fmt.Sprintf("%v [INFO] dummy task duration %d", time.Now(), randDuration))
	dummyOp.EndedAt = time.Now()

	return nil
}

// GetLogs ..
func (dummyOp *DummyOperator) GetLogs() []string {
	return dummyOp.Logs
}
