package osrm

import geo "github.com/paulmach/go.geo"

// RouteRequest represents a request to the route method
type RouteRequest struct {
	Request
	Bearings    []Bearing
	Steps       Steps
	Annotations Annotations
	Overview    Overview
	Geometries  Geometries
}

// RouteResponse represents a response from the route method
type RouteResponse struct {
	ResponseError
	Routes []Route `json:"routes"`
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

func (r RouteRequest) request() *Request {
	r.service = "route"
	r.options = stepsOptions(r.Steps, r.Annotations, r.Overview, r.Geometries)
	r.options.Set("continue_straight", "true")
	if len(r.Bearings) > 0 {
		r.options.Set("bearings", bearings(r.Bearings))
	}
	return &r.Request
}

func stepsOptions(s Steps, a Annotations, o Overview, g Geometries) Options {
	options := Options{}

	if steps := s.String(); steps != "" {
		options.Set("steps", steps)
	}

	if annotations := a.String(); annotations != "" {
		options.Set("annotations", annotations)
	}

	options.Set("geometries", GeometriesPolyline6.String())
	if geometries := g.String(); geometries != "" {
		options.Set("geometries", geometries)
	}

	if overview := o.String(); overview != "" {
		options.Set("overview", overview)
	}

	return options
}
