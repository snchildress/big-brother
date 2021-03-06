package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/jmoiron/jsonq"
)

// getEstimate requests an estimated price for the configured
// coordinates for the given service, Lyft or Uber
func getEstimate(lyft bool, startAddress address, endAddress address) (estimate string, err error) {
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
		queryString.Add("start_lat", fmt.Sprintf("%f", startAddress.Latitude))
		queryString.Add("start_lng", fmt.Sprintf("%f", startAddress.Longitude))
		queryString.Add("end_lat", fmt.Sprintf("%f", endAddress.Latitude))
		queryString.Add("end_lng", fmt.Sprintf("%f", endAddress.Longitude))
	} else {
		req.Header.Set("Authorization", "Token "+uberAPIKey)
		queryString.Add("start_latitude", fmt.Sprintf("%f", startAddress.Latitude))
		queryString.Add("start_longitude", fmt.Sprintf("%f", startAddress.Longitude))
		queryString.Add("end_latitude", fmt.Sprintf("%f", endAddress.Latitude))
		queryString.Add("end_longitude", fmt.Sprintf("%f", endAddress.Longitude))
	}
	req.URL.RawQuery = queryString.Encode()

	// Make the API request
	res, err := client.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()

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
		estimate = "Lyft: $" + lyftMinPrice + "-" + lyftMaxPrice
	} else {
		uberEstimate, _ := resBody.String("prices", "0", "estimate")
		estimate = "Uber: " + uberEstimate
	}
	return
}
