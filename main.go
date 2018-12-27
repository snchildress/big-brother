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

// Handler is the main AWS Lambda handler func
func Handler(request events.APIGatewayProxyRequest) (response events.APIGatewayProxyResponse, err error) {
	// GET the Lyft estimated cost
	client := &http.Client{}
	req, err := http.NewRequest("GET", lyftEndpoint, nil)
	if err != nil {
		return
	}
	queryString := req.URL.Query()
	queryString.Add("start_lat", startLat)
	queryString.Add("start_lng", startLong)
	queryString.Add("end_lat", endLat)
	queryString.Add("end_lng", endLong)
	req.URL.RawQuery = queryString.Encode()
	req.Header.Set("Authorization", "bearer "+lyftAPIKey)
	res, err := client.Do(req)
	if err != nil {
		return
	}
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

	response = events.APIGatewayProxyResponse{
		Body:       lyftEstimateMessage,
		StatusCode: 200,
	}
	return
}

func main() {
	lambda.Start(Handler)
}
