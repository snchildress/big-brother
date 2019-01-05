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
func getAddress(addresses []address) (address address, err error) {
	rand.Seed(time.Now().Unix())           // resets randomization on each function call
	randomInt := rand.Intn(len(addresses)) // get random item in addresses array
	address = addresses[randomInt]
	return
}
