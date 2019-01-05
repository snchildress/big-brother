package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Handler is the main AWS Lambda handler func
func Handler() (response events.APIGatewayProxyResponse, err error) {
	// Get random addresses for origin and destination
	addresses, err := getAddresses()
	if err != nil {
		return
	}
	startAddress, addresses, err := getAddress(addresses)
	if err != nil {
		return
	}
	endAddress, addresses, err := getAddress(addresses)
	if err != nil {
		return
	}

	originAddress := fmt.Sprintf("%v", startAddress.Name) + " (" + startAddress.Neighborhood + ")"
	destinationAddress := fmt.Sprintf("%v", endAddress.Name) + " (" + endAddress.Neighborhood + ")"

	// Get the Lyft estimate
	lyftEstimate, err := getEstimate(true, startAddress, endAddress)
	if err != nil {
		return
	}

	// Get the Uber estimate
	uberEstimate, err := getEstimate(false, startAddress, endAddress)
	if err != nil {
		return
	}

	// Return the estimated prices in the response body
	headers := map[string]string{"Content-Type": "text/plain"}
	message := "Current estimated New Orleans rideshare prices:" +
		"\n\nOrigin: " + originAddress + "\nDestination: " +
		destinationAddress + "\n\n" + lyftEstimate + "\n" + uberEstimate
	response = events.APIGatewayProxyResponse{
		Headers:    headers,
		Body:       message,
		StatusCode: 200,
	}
	return
}

func main() {
	lambda.Start(Handler)
}
