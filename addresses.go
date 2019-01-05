package main

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

type address struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Address   int     `json:"address"`
	Street    string  `json:"street"`
}

// getAddresses converts the addresses.json blob to an
// instance of the addresses struct
func getAddresses() (addresses []address, err error) {
	jsonFile, err := os.Open("addresses.json")
	if err != nil {
		return
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal([]byte(byteValue), &addresses)
	return
}

// getAddress gets a random address
func getAddress(addresses []address) (address address, returnedAddresses []address, err error) {
	rand.Seed(time.Now().Unix())             // resets randomization on each function call
	randomIndex := rand.Intn(len(addresses)) // get random item in addresses array
	address = addresses[randomIndex]

	// Remove the chosen address
	addresses[randomIndex] = addresses[len(addresses)-1]
	nilAddress := address
	addresses[len(addresses)-1] = nilAddress
	returnedAddresses = addresses[:len(addresses)-1]
	return
}
