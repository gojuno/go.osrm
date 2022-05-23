package osrm

// TableRequest represents a request to the table method
type TableRequest struct {
	Profile               string
	Coordinates           Geometry
	Sources, Destinations []int
	Annotations           []string
}

// TableResponse resresents a response from the table method
type TableResponse struct {
	ResponseStatus
	Durations [][]float32 `json:"durations"`
}

func (r TableRequest) request() *request {
	opts := options{}
	if len(r.Sources) > 0 {
		opts.addInt("sources", r.Sources...)
	}
	if len(r.Destinations) > 0 {
		opts.addInt("destinations", r.Destinations...)
	}
	if len(r.Annotations) > 0 {
		opts.add("annotations", r.Annotations...)
	}

	return &request{
		profile: r.Profile,
		coords:  r.Coordinates,
		service: "table",
		options: opts,
	}
}
