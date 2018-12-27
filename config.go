package main

import "os"

// API keys
var uberAPIKey = os.Getenv("UBER_API_KEY")
var lyftAPIKey = os.Getenv("LYFT_API_KEY")
var slackAPIKey = os.Getenv("SLACK_API_KEY")

// API Endpoint
var uberEndpoint = "https://api.uber.com/v1.2/estimates/price"
var lyftEndpoint = "https://api.lyft.com/v1/cost"

// Trip origin and destination coordinates

// 365 Canal St, New Orleans, LA 70130
var startLat = "29.951230"
var startLong = "-90.065490"

// Finn McCool's Irish Pub, Banks St, New Orleans, LA 70119
var endLat = "29.969600"
var endLong = "-90.099240"
