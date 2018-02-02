package osrm

import geo "github.com/paulmach/go.geo"

// MatchRequest represents a request to the match method
type MatchRequest struct {
	Request
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
	ResponseError
	Matchings   []Matching  `json:"matchings"`
	Tracepoints []*Waypoint `json:"tracepoints"`
}

// Matching represents an array of Route objects that assemble the trace
type Matching struct {
	Route
	Confidence float64 `json:"confidence"`
	Geometry   GeoPath `json:"geometry"`
}

func (r MatchRequest) request() *Request {
	r.service = "match"
	r.options = matcherOptions(
		stepsOptions(r.Steps, r.Annotations, r.Overview, r.Geometries),
		r.Tidy,
		r.Gaps,
	)

	if len(r.Timestamps) > 0 {
		r.options.AddInt64("timestamps", r.Timestamps...)
	}
	if len(r.Radiuses) > 0 {
		r.options.AddFloat("radiuses", r.Radiuses...)
	}
	if len(r.Hints) > 0 {
		r.options.Add("hints", r.Hints...)
	}

	return &r.Request
}

// Waypoint represents a matched point on a route
type Waypoint struct {
	Index             int       `json:"waypoint_index"`
	Location          geo.Point `json:"location"`
	MatchingIndex     int       `json:"matchings_index"`
	AlternativesCount int       `json:"alternatives_count"`
	Hint              string    `json:"hint"`
}

func matcherOptions(options Options, t Tidy, g Gaps) Options {
	if tidy := t.String(); tidy != "" {
		options.Set("tidy", tidy)
	}

	if gaps := g.String(); gaps != "" {
		options.Set("gaps", gaps)
	}
	return options
}
