package main

import "time"

// Advice - traffic advice
type Advice struct {
	InTraffic            bool      `json:"inTraffic"`
	Speed                float64   `json:"speed"`
	CarLengths           float64   `json:"carLengths"`
	TrafficDistanceStart float64   `json:"trafficDistanceStart"`
	TrafficDistanceEnd   float64   `json:"trafficDistanceEnd"`
	SpeedLimit           int       `json:"speedLimit"`
	Timestamp            time.Time `json:"timestamp"`
}
