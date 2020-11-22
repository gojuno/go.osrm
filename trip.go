package osrm

type TripRequest struct {
	Profile     string
	Coordinates Geometry
	Roundtrip   Roundtrip
	Source      Source
	Destination Destination
	Steps       Steps
	Annotations Annotations
	Geometries  Geometries
	Overview    Overview
}

type TripResponse struct {
	ResponseStatus
	Waypoints []TripWaypoint `json:"waypoints"`
	Trips     []Route        `json:"trips"`
}

type TripWaypoint struct {
	TripsIndex    int `json:"trips_index"`
	WaypointIndex int `json:"waypoint_index"`
	Waypoint
}

func (r TripRequest) request() *request {
	return &request{
		profile: r.Profile,
		coords:  r.Coordinates,
		service: "trip",
		options: stepsOptions(r.Steps, r.Annotations, r.Overview, r.Geometries).
			setStringer("roundtrip", valueOrDefault(r.Roundtrip, RoundtripDefault)).
			setStringer("source", valueOrDefault(r.Source, SourceDefault)).
			setStringer("destination", valueOrDefault(r.Destination, DestinationDefault)),
	}
}

func (r TripRequest) IsSupported() bool {
	fixedstart := r.Source == SourceFirst || (r.Source == SourceAny && r.Destination == DestinationAny)
	fixedend := r.Destination == DestinationLast
	roundtrip := r.Roundtrip == RoundtripTrue
	if fixedstart && fixedend && !roundtrip {
		return true
	} else if roundtrip {
		return true
	}
	return false
}
