package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Response is a JSON response object
type Response struct {
	Message string `json:"message"`
}

// Handler is the main AWS Lambda handler func
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		Body:       "Hello, World!",
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
