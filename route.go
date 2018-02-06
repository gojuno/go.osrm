package osrm

import geo "github.com/paulmach/go.geo"

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
	opts := stepsOptions(r.Steps, r.Annotations, r.Overview, r.Geometries)

	if cs := r.ContinueStraight.String(); cs != "" {
		opts.set("continue_straight", cs)
	}

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

func stepsOptions(s Steps, a Annotations, o Overview, g Geometries) options {
	opts := options{}

	if steps := s.String(); steps != "" {
		opts.set("steps", steps)
	}

	if annotations := a.String(); annotations != "" {
		opts.set("annotations", annotations)
	}

	opts.set("geometries", GeometriesPolyline6.String())
	if geometries := g.String(); geometries != "" {
		opts.set("geometries", geometries)
	}

	if overview := o.String(); overview != "" {
		opts.set("overview", overview)
	}

	return opts
}
