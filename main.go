package main

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// ProxyHandler handles the incoming Lambda event and forwards the HTTP request
func ProxyHandler(context.Context, *events.LambdaFunctionURLRequest) (*events.LambdaFunctionURLResponse, error) {
	res := map[string]interface{}{
		"hello": "world",
	}

	body, err := json.Marshal(res)
	if err != nil {
		return nil, err
	}

	return &events.LambdaFunctionURLResponse{
		StatusCode: 200,
		Headers:    map[string]string{"Content-Type": "application/json"},
		Body:       string(body),
	}, nil
}

func main() {
	// Make the handler available for Lambda
	lambda.StartHandlerFunc(ProxyHandler)
}
