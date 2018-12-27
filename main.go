package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Response is a JSON response object
type Response struct {
	Message string `json:"message"`
}

// Handler is the main AWS Lambda handler func
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// GET the Lyft estimated cost
	client := &http.Client{}
	req, _ := http.NewRequest("GET", lyftEndpoint, nil)
	queryString := req.URL.Query()
	queryString.Add("start_lat", startLat)
	queryString.Add("start_lng", startLong)
	queryString.Add("end_lat", endLat)
	queryString.Add("end_lng", endLong)
	req.URL.RawQuery = queryString.Encode()
	req.Header.Set("Authorization", "bearer "+lyftAPIKey)
	res, _ := client.Do(req)
	resBodyBytes, _ := ioutil.ReadAll(res.Body)
	resBody := string(resBodyBytes)
	fmt.Println(resBody)

	return events.APIGatewayProxyResponse{
		Body:       "Hello, World!",
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
