package osrm

import geo "github.com/paulmach/go.geo"

// MatchRequest represents a request to the match method
type MatchRequest struct {
	Profile     string
	Coordinates Geometry
	Bearings    []Bearing
	Steps       Steps
	Annotations Annotations
	Tidy        Tidy
	Timestamps  []int64
	Radiuses    []float64
	Hints       []string
	Overview    Overview
	Gaps        Gaps
	Geometries  Geometries
}

// MatchResponse represents a response from the match method
type MatchResponse struct {
	Matchings   []Matching    `json:"matchings"`
	Tracepoints []*Tracepoint `json:"tracepoints"`
}

type matchResponseOrError struct {
	responseStatus
	MatchResponse
}

// Matching represents an array of Route objects that assemble the trace
type Matching struct {
	Route
	Confidence float64  `json:"confidence"`
	Geometry   Geometry `json:"geometry"`
}

func (r MatchRequest) request() *request {
	options := matcherOptions(
		stepsOptions(r.Steps, r.Annotations, r.Overview, r.Geometries),
		r.Tidy,
		r.Gaps,
	)
	if len(r.Timestamps) > 0 {
		options.addInt64("timestamps", r.Timestamps...)
	}
	if len(r.Radiuses) > 0 {
		options.addFloat("radiuses", r.Radiuses...)
	}
	if len(r.Hints) > 0 {
		options.add("hints", r.Hints...)
	}
	if len(r.Bearings) > 0 {
		options.set("bearings", bearings(r.Bearings))
	}

	return &request{
		profile: r.Profile,
		coords:  r.Coordinates,
		service: "match",
		options: options,
	}
}

// URL generates a url for OSRM match request
func (r MatchRequest) URL(serverURL string) (string, error) {
	return r.request().URL(serverURL)
}

// Tracepoint represents a matched point on a route
type Tracepoint struct {
	Index             int       `json:"waypoint_index"`
	Location          geo.Point `json:"location"`
	MatchingIndex     int       `json:"matchings_index"`
	AlternativesCount int       `json:"alternatives_count"`
	Hint              string    `json:"hint"`
}

func matcherOptions(options options, tidy Tidy, gaps Gaps) options {
	return options.
		setStringer("tidy", tidy).
		setStringer("gaps", gaps)
}
