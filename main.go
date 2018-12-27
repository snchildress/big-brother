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
func Handler() (response events.APIGatewayProxyResponse, err error) {
	// Get the Lyft estimate
	lyftEstimate, err := getEstimate(true)
	if err != nil {
		return
	}

	// Get the Uber estimate
	uberEstimate, err := getEstimate(false)
	if err != nil {
		return
	}

	// Print out the two services' estimates
	fmt.Println(lyftEstimate)
	fmt.Println(uberEstimate)

	// Return the Lyft estimate
	response = events.APIGatewayProxyResponse{
		Body:       lyftEstimate,
		StatusCode: 200,
	}
	return
}

// getEstimate requests an estimated price for the configured
// coordinates for the given service, Lyft or Uber
func getEstimate(lyft bool) (estimate string, err error) {
	// Create an HTTP client for the given service's endpoint and key
	client := &http.Client{}
	endpoint := uberEndpoint
	if lyft {
		endpoint = lyftEndpoint
	}
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return
	}

	// Set the headers and query string accordingly
	queryString := req.URL.Query()
	if lyft {
		req.Header.Set("Authorization", "bearer "+lyftAPIKey)
		queryString.Add("start_lat", startLat)
		queryString.Add("start_lng", startLong)
		queryString.Add("end_lat", endLat)
		queryString.Add("end_lng", endLong)
	} else {
		req.Header.Set("Authorization", "Token "+uberAPIKey)
		queryString.Add("start_latitude", startLat)
		queryString.Add("start_longitude", startLong)
		queryString.Add("end_latitude", endLat)
		queryString.Add("end_longitude", endLong)
	}
	req.URL.RawQuery = queryString.Encode()

	// Make the API request
	res, err := client.Do(req)
	if err != nil {
		return
	}

	// Get the response body
	resBodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	resBodyString := string(resBodyBytes)
	resBodyMap := map[string]interface{}{}
	json.NewDecoder(strings.NewReader(resBodyString)).Decode(&resBodyMap)
	resBody := jsonq.NewQuery(resBodyMap)

	// Get the estimate from the response body
	if lyft {
		lyftEstimate, _ := resBody.Object("cost_estimates", "0")
		lyftMinPriceFloat := lyftEstimate["estimated_cost_cents_min"].(float64) / 100
		lyftMaxPriceFloat := lyftEstimate["estimated_cost_cents_max"].(float64) / 100
		lyftMinPrice := fmt.Sprintf("%.0f", lyftMinPriceFloat)
		lyftMaxPrice := fmt.Sprintf("%.0f", lyftMaxPriceFloat)
		estimate = "Lyft Estimate: $" + lyftMinPrice + "-" + lyftMaxPrice
	} else {
		uberEstimate, _ := resBody.String("prices", "0", "estimate")
		estimate = "Uber Estimate: " + uberEstimate
	}
	return
}

func main() {
	lambda.Start(Handler)
}
