package osrm

// TableRequest represents a request to the table method
type TableRequest struct {
	GeneralOptions
	Profile               string
	Coordinates           Geometry
	Sources, Destinations []int
	Annotations           Annotations
	FallbackSpeed         float64
	FallbackCoordinate    FallbackCoordinate
	ScaleFactor           float64
}

// TableResponse resresents a response from the table method
type TableResponse struct {
	ResponseStatus
	Durations          [][]float32 `json:"durations"`
	Distances          [][]float32 `json:"distances"`
	Sources            []Waypoint  `json:"sources"`
	Destinations       []Waypoint  `json:"destinations"`
	FallbackSpeedCells [][]int     `json:"fallback_speed_cells"`
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
		opts.setStringer("annotations", r.Annotations)
	}
	if r.FallbackSpeed > 0 {
		opts.addFloat("fallback_speed", r.FallbackSpeed)
	}
	if r.FallbackCoordinate.Valid() {
		opts.setStringer("fallback_coordinate", r.FallbackCoordinate)
	}
	if r.ScaleFactor > 0 {
		opts.addFloat("scale_factor", r.ScaleFactor)
	}

	opts = r.GeneralOptions.options(opts)

	return &request{
		profile: r.Profile,
		coords:  r.Coordinates,
		service: "table",
		options: opts,
	}
}
