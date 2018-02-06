package osrm

import (
	"fmt"

	geo "github.com/paulmach/go.geo"
)

// RouteRequest represents a request to the route method
type RouteRequest struct {
	Profile          string
	GeoPath          GeoPath
	Bearings         []Bearing
	Steps            Steps
	Annotations      Annotations
	Overview         Overview
	Geometries       Geometries
	ContinueStraight ContinueStraight
}

// RouteResponse represents a response from the route method
type RouteResponse struct {
	Routes []Route `json:"routes"`
}

type routeResponseOrError struct {
	responseStatus
	RouteResponse
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
}

// RouteStep represents a route geometry
type RouteStep struct {
	Distance float32      `json:"distance"`
	Duration float32      `json:"duration"`
	Name     string       `json:"name"`
	Geometry GeoPath      `json:"geometry"`
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

	if len(r.Bearings) > 0 {
		opts.set("bearings", bearings(r.Bearings))
	}

	return &request{
		profile: r.Profile,
		geoPath: r.GeoPath,
		service: "route",
		options: opts,
	}
}

// URL generates a url for OSRM route request
func (r RouteRequest) URL(serverURL string) (string, error) {
	return r.request().URL(serverURL)
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
