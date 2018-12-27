package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type addresses struct {
	Address []address `json:"addresses"`
}

type address struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Address   int     `json:"address"`
	Street    string  `json:"street"`
}

// getAddresses converts the addresses.json blob to an
// instance of the addresses struct
func getAddresses() (addresses addresses, err error) {
	jsonFile, err := os.Open("addresses.json")
	if err != nil {
		return
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal([]byte(byteValue), &addresses)
	return
}
