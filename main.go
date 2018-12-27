package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/jmoiron/jsonq"
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
	data := map[string]interface{}{}
	dec := json.NewDecoder(strings.NewReader(resBody))
	dec.Decode(&data)
	jq := jsonq.NewQuery(data)
	lyftEstimate, _ := jq.Object("cost_estimates", "0")
	lyftMinPriceFloat := lyftEstimate["estimated_cost_cents_min"].(float64) / 100
	lyftMaxPriceFloat := lyftEstimate["estimated_cost_cents_max"].(float64) / 100
	lyftMinPrice := fmt.Sprintf("%.0f", lyftMinPriceFloat)
	lyftMaxPrice := fmt.Sprintf("%.0f", lyftMaxPriceFloat)
	lyftEstimateMessage := "Lyft Estimate: $" + lyftMinPrice + "-" + lyftMaxPrice

	return events.APIGatewayProxyResponse{
		Body:       lyftEstimateMessage,
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
