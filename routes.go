package main

import "net/http"

// Route - a method of routing a REST request
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes - all of the routes we route
type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"NewSession",
		"POST",
		"/new",
		NewSession,
	},
	Route{
		"NewGeoState",
		"POST",
		"/geostate",
		NewGeoState,
	},
}
