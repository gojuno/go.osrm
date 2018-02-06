package osrm

// TableRequest represents a request to the table method
type TableRequest struct {
	Profile               string
	GeoPath               GeoPath
	Sources, Destinations []int
}

// TableResponse resresents a response from the table method
type TableResponse struct {
	Durations [][]float32 `json:"durations"`
}

type tableResponseOrError struct {
	responseStatus
	TableResponse
}

func (r TableRequest) request() *request {
	opts := options{}
	if len(r.Sources) > 0 {
		opts.addInt("sources", r.Sources...)
	}
	if len(r.Destinations) > 0 {
		opts.addInt("destinations", r.Destinations...)
	}

	return &request{
		profile: r.Profile,
		geoPath: r.GeoPath,
		service: "table",
		options: opts,
	}
}

// URL generates a url for OSRM table request
func (r TableRequest) URL(serverURL string) (string, error) {
	return r.request().URL(serverURL)
}
