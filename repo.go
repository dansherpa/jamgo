package main

import (
	"fmt"
	"time"

	"github.com/nu7hatch/gouuid"
)

var sessions Sessions

// RepoFindSession finds an existing session, returns nil if not found.
func RepoFindSession(uuidString string) (session Session, err error) {
	for _, s := range sessions {
		if s.UUID == uuidString {
			return s, nil
		}
	}

	return session, fmt.Errorf("Could not find Session with uuid of %s", uuidString)
}

// RepoCreateSession generates a new session.
func RepoCreateSession() (Session, error) {
	var s Session
	var u *uuid.UUID
	var err error
	u, err = uuid.NewV4()
	if err != nil {
		fmt.Println("error:", err)
		return s, err
	}
	s.UUID = u.String()
	s.Timestamp = time.Now()
	sessions = append(sessions, s)
	return s, nil
}

// RepoDestroySession destroys an existing session. Returns an error if not found.
func RepoDestroySession(uuidString string) error {
	for i, s := range sessions {
		if s.UUID == uuidString {
			sessions = append(sessions[:i], sessions[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Session with uuid of %s to delete", uuidString)
}

// RepoNewGeoState reports a new GeoState for a session.
func RepoNewGeoState(geo GeoState) Advice {
	var adv Advice

	// for now just make shit up
	// in the future, log the geostate here
	// calculate the advice
	adv.CarLengths = 7
	adv.Speed = geo.Speed / 2.0
	adv.Timestamp = time.Now()
	adv.TrafficDistanceStart = 5.01
	adv.TrafficDistanceEnd = 6.02
	adv.SpeedLimit = 65

	return adv
}
