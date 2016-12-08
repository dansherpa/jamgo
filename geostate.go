package main

import "time"

// GeoState - the geographic state at a given time.
type GeoState struct {
	SessionUUID string    `json:"uuid"`
	Speed       float64   `json:"speed"`
	Heading     float64   `json:"heading"`
	Latitude    float64   `json:"latitude"`
	Longitude   float64   `json:"longitude"`
	Timestamp   time.Time `json:"timestamp"`
}

// GeoStates - Collection of geo states.
type GeoStates []GeoState
