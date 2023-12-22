package osrm

import (
	"strconv"
)

// TripRequest represents a request to the trip method
type TripRequest struct {
	Profile     string
	Coordinates Geometry
	Bearings    []Bearing
	Steps       Steps
	Annotations Annotations
	Overview    Overview
	Geometries  Geometries
	Roundtrip   Roundtrip
	Source      Source
	Destination Destination
	Waypoints   []int
}

// TripResponse represents a response from the trip method
type TripResponse struct {
	ResponseStatus
	Trips     []Route        `json:"trips"`
	Waypoints []TripWaypoint `json:"waypoints"`
}

func (r TripRequest) request() *request {
	opts := stepsOptions(r.Steps, r.Annotations, r.Overview, r.Geometries).
		setStringer("source", r.Source).
		setStringer("destination", r.Destination).
		setStringer("roundtrip", r.Roundtrip)

	if len(r.Waypoints) > 0 {
		waypoints := ""
		for i, w := range r.Waypoints {
			if i > 0 {
				waypoints += ";"
			}
			waypoints += strconv.Itoa(w)
		}
		opts.set("waypoints", waypoints)
	}

	if len(r.Bearings) > 0 {
		opts.set("bearings", bearings(r.Bearings))
	}

	return &request{
		profile: r.Profile,
		coords:  r.Coordinates,
		service: "trip",
		options: opts,
	}
}

// TripWaypoint represent a waypoint in a trip
type TripWaypoint struct {
	Waypoint
	Index      int `json:"waypoint_index"`
	TripsIndex int `json:"trips_index"`
}
