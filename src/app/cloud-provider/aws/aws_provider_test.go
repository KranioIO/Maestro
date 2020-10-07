package aws_test

import (
	"app/cloud-provider/aws"
	"log"
	"testing"
)

func TestCallLambdaNoPayload(t *testing.T) {
	result, err := aws.CallLambda("test-lambda-dev-hello", nil)

	if err != nil {
		t.Error(err)
	} else {
		log.Println(result)
	}
}

func TestCallLambdaWithPayload(t *testing.T) {
	result, err := aws.CallLambda("test-lambda-dev-hello", struct{ Data string }{Data: "data-value"})

	if err != nil {
		t.Error(err)
	} else {
		log.Println(result)
	}
}
