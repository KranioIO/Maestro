package aws

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
)

const defaultRegion = "us-east-1"

// CallLambda invokes an lambda by name synchronously
func CallLambda(name string, payload interface{}) (map[string]interface{}, error) {
	payloadBytes, err := json.Marshal(payload)

	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	ctx, cancelFn := context.WithTimeout(ctx, 900)

	if cancelFn != nil {
		defer cancelFn()
	}

	sess := session.Must(session.NewSession())
	lambdaClient := lambda.New(sess, &aws.Config{Region: aws.String(defaultRegion)})

	response, err := lambdaClient.Invoke(&lambda.InvokeInput{
		FunctionName: aws.String(name),
		Payload:      payloadBytes,
	})

	if err != nil {
		return nil, err
	}

	return parseLambdaResponse(response)
}

func parseLambdaResponse(response *lambda.InvokeOutput) (map[string]interface{}, error) {
	var data interface{}

	if status := *response.StatusCode; status != 200 {
		return nil, errors.New("lambda response: desired status code 200, received " + string(status))
	}

	if err := json.Unmarshal(response.Payload, &data); err != nil {
		return nil, err
	}

	return data.(map[string]interface{}), nil
}
