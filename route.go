package osrm

import (
	"fmt"

	geo "github.com/paulmach/go.geo"
)

// RouteRequest represents a request to the route method
type RouteRequest struct {
	GeneralOptions
	Profile          string
	Coordinates      Geometry
	Steps            Steps
	Annotations      Annotations
	Overview         Overview
	Geometries       Geometries
	ContinueStraight ContinueStraight
}

// RouteResponse represents a response from the route method
type RouteResponse struct {
	ResponseStatus
	Routes    []Route    `json:"routes"`
	Waypoints []Waypoint `json:"waypoints"`
}

// Route represents a route through (potentially multiple) points.
type Route struct {
	Distance float32    `json:"distance"`
	Duration float32    `json:"duration"`
	Legs     []RouteLeg `json:"legs"`
}

// RouteLeg represents a route between two waypoints.
type RouteLeg struct {
	Annotation Annotation  `json:"annotation"`
	Distance   float32     `json:"distance"`
	Duration   float32     `json:"duration"`
	Steps      []RouteStep `json:"steps"`
}

// Annotation contains additional metadata for each coordinate along the route geometry
type Annotation struct {
	Duration []float32 `json:"duration"`
	Distance []float32 `json:"distance"`
	Nodes    []uint64  `json:"nodes"`
}

// RouteStep represents a route geometry
type RouteStep struct {
	Distance float32      `json:"distance"`
	Duration float32      `json:"duration"`
	Name     string       `json:"name"`
	Geometry Geometry     `json:"geometry"`
	Mode     string       `json:"mode"`
	Maneuver StepManeuver `json:"maneuver"`
}

// StepManeuver contains information about maneuver in step
type StepManeuver struct {
	BearingBefore float32   `json:"bearing_before"`
	BearingAfter  float32   `json:"bearing_after"`
	Location      geo.Point `json:"location"`
	Type          string    `json:"type"`
	Modifier      string    `json:"modifier"`
}

func (r RouteRequest) request() *request {
	opts := stepsOptions(r.Steps, r.Annotations, r.Overview, r.Geometries).
		setStringer("continue_straight", r.ContinueStraight)

	opts = r.GeneralOptions.options(opts)

	return &request{
		profile: r.Profile,
		coords:  r.Coordinates,
		service: "route",
		options: opts,
	}
}

func stepsOptions(steps Steps, annotations Annotations, overview Overview, geometries Geometries) options {
	return options{}.
		setStringer("steps", steps).
		setStringer("annotations", annotations).
		setStringer("geometries", valueOrDefault(geometries, GeometriesPolyline6)).
		setStringer("overview", overview)
}

func valueOrDefault(value, def fmt.Stringer) fmt.Stringer {
	if value.String() == "" {
		return def
	}
	return value
}
