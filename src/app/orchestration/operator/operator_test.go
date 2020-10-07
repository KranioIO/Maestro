package operator_test

import (
	"app/orchestration/operator"
	"log"
	"testing"
)

func TestDummyOperatorExecution(t *testing.T) {
	dummy := operator.DummyOperator{}
	err := dummy.Execute()

	if err == nil {
		log.Println("dummy operator started at", dummy.StartedAt)
		log.Println("dummy operator ended at", dummy.EndedAt)
		log.Println("dummy operator logs", dummy.Logs)
	}
}

func TestLambdaOperatorExecution(t *testing.T) {
	lambda := operator.LambdaOperator{
		FunctionPayload: struct{ Data string }{Data: "test data value"},
		FunctionName:    "test-lambda-dev-hello",
	}

	if err := lambda.Execute(); err != nil {
		t.Error(err)
	}

	log.Println("lambda operator started at", lambda.StartedAt)
	log.Println("lambda operator ended at", lambda.EndedAt)
	log.Println("lambda operator result", lambda.FunctionResult)
	log.Println("lambda operator logs", lambda.Logs)
}
