package osrm

import (
	"fmt"

	geo "github.com/paulmach/go.geo"
)

// RouteRequest represents a request to the route method
type RouteRequest struct {
	Profile          string
	Coordinates      Geometry
	Bearings         []Bearing
	Steps            Steps
	Annotations      Annotations
	Overview         Overview
	Geometries       Geometries
	ContinueStraight ContinueStraight
}

// RouteResponse represents a response from the route method
type RouteResponse struct {
	ResponseStatus
	Routes []Route `json:"routes"`
}

// Route represents a route through (potentially multiple) points.
type Route struct {
	Distance   float32    `json:"distance"`
	Duration   float32    `json:"duration"`
	WeightName string     `json:"weight_name"`
	Wieght     float32    `json:"weight"`
	Geometry   Geometry   `json:"geometry"`
	Legs       []RouteLeg `json:"legs"`
}

// RouteLeg represents a route between two waypoints.
type RouteLeg struct {
	Annotation Annotation  `json:"annotation"`
	Distance   float32     `json:"distance"`
	Duration   float32     `json:"duration"`
	Summary    string      `json:"summary"`
	Weight     float32     `json:"weight"`
	Steps      []RouteStep `json:"steps"`
}

// Annotation contains additional metadata for each coordinate along the route geometry
type Annotation struct {
	Duration []float32 `json:"duration,omitempty"`
	Distance []float32 `json:"distance,omitempty"`
	Nodes    []uint64  `json:"nodes,omitempty"`
}

// RouteStep represents a route geometry
type RouteStep struct {
	Distance      float32        `json:"distance"`
	Duration      float32        `json:"duration"`
	Geometry      Geometry       `json:"geometry"`
	Name          string         `json:"name"`
	Mode          string         `json:"mode"`
	DrivingSide   string         `json:"driving_side"`
	Weight        float32        `json:"weight"`
	Maneuver      StepManeuver   `json:"maneuver"`
	Intersections []Intersection `json:"intersections,omitempty"`
}

type Intersection struct {
	Location geo.Point `json:"location"`
	Bearings []uint16  `json:"bearings"`
	Entry    []bool    `json:"entry"`
	In       *uint32   `json:"in,omitempty"`
	Out      *uint32   `json:"out,omitempty"`
	Lanes    []Lane    `json:"lanes,omitempty"`
}

type Lane struct {
	Indications []string `json:"indications"`
	Valid       bool     `json:"valid"`
}

// StepManeuver contains information about maneuver in step
type StepManeuver struct {
	Location      geo.Point `json:"location"`
	BearingBefore float32   `json:"bearing_before"`
	BearingAfter  float32   `json:"bearing_after"`
	Type          string    `json:"type"`
	Modifier      string    `json:"modifier,omitempty"`
	Exit          *uint32   `json:"exit,omitempty"`
}

func (r RouteRequest) request() *request {
	opts := stepsOptions(r.Steps, r.Annotations, r.Overview, r.Geometries).
		setStringer("continue_straight", r.ContinueStraight)

	if len(r.Bearings) > 0 {
		opts.set("bearings", bearings(r.Bearings))
	}

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
