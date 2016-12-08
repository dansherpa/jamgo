package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// Index - just welcome
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}

// NewSession creates a new session.
/*
Test with this curl command:

curl -H "Content-Type: application/json" -d '{}' http://localhost:8080/new

*/
func NewSession(w http.ResponseWriter, r *http.Request) {
	/*var todo Todo
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &todo); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}*/
	var err error
	var s Session
	s, err = RepoCreateSession()
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err = json.NewEncoder(w).Encode(s); err != nil {
		panic(err)
	}
}

// NewGeoState uploads a new geostate and downloads advice
/*
Test with this curl command:

curl -H "Content-Type: application/json" -d '{"uuid":"foo","speed":30,"heading":360,"latitude":41.454334, "longitude":-120.5467887, "timestamp":"2016-12-08T14:51:52.336018393-05:00"}' http://localhost:8888/geostate

*/
func NewGeoState(w http.ResponseWriter, r *http.Request) {
	var geo GeoState
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	//	fmt.Printf("body: %s\n", body)
	if err := json.Unmarshal(body, &geo); err != nil {
		fmt.Print("got an error\n")
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err = json.NewEncoder(w).Encode(err); err != nil {
			//			fmt.Printf("Encoding error: %+v\n", err)
			panic(err)
		}
		//		fmt.Printf("Unmarshall error: %+v\n", err.Error())
		panic(err)
	}

	//	fmt.Printf("geo: %+v\n", geo)
	var advice Advice
	advice = RepoNewGeoState(geo)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(advice); err != nil {
		panic(err)
	}
}
