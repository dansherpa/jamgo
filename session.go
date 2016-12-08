package main

import "time"

// Session - a client session
type Session struct {
	UUID      string    `json:"uuid"`
	Timestamp time.Time `json:"timestamp"`
}

// Sessions - a collection of sessions
type Sessions []Session
