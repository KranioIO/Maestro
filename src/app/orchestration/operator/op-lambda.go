package operator

import (
	"app/cloud-provider/aws"
	"time"
)

// LambdaOperator calls a lambda function synchronously
type LambdaOperator struct {
	FunctionName    string
	FunctionPayload interface{}
	FunctionResult  map[string]interface{}
	StartedAt       time.Time
	EndedAt         time.Time
	Logs            []string
}

// CreateLambda returns a dummy operator instance
func CreateLambda(payload map[interface{}]interface{}) Operator {
	return &LambdaOperator{
		FunctionName: payload["name"].(string),
	}
}

// Execute will perform the goal of the operator
func (lambdaOP *LambdaOperator) Execute() error {
	lambdaOP.StartedAt = time.Now()
	result, err := aws.CallLambda(lambdaOP.FunctionName, lambdaOP.FunctionPayload)

	if err != nil {
		lambdaOP.Logs = append(lambdaOP.Logs, err.Error())
	}

	for _, log := range result { // TODO: improve log save
		if log != nil {
			lambdaOP.Logs = append(lambdaOP.Logs, log.(string))
		}
	}

	lambdaOP.FunctionResult = result
	lambdaOP.EndedAt = time.Now()

	return err
}

// GetLogs ..
func (lambdaOP *LambdaOperator) GetLogs() []string {
	return lambdaOP.Logs
}
