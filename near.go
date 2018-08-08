package osrm

// NearRequest represents a request to the near method
type NearRequest struct {
	Profile  string
	GeoPath  GeoPath
	Bearings []Bearing
	Number   int
}

// NearResponse represents a response from the near method
type NearResponse struct {
	Waypoints []Waypoint `json:"waypoints"`
}

type nearResponseOrError struct {
	responseStatus
	NearResponse
}

func (r NearRequest) request() *request {
	opts := options{}
	opts.addInt("number", r.Number)

	if len(r.Bearings) > 0 {
		opts.set("bearings", bearings(r.Bearings))
	}

	return &request{
		profile: r.Profile,
		service: "near",
		options: opts,
	}
}

// URL generates a url for OSRM match request
func (r NearRequest) URL(serverURL string) (string, error) {
	return r.request().URL(serverURL)
}
