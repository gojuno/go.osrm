package osrm

import geo "github.com/paulmach/go.geo"

// NearRequest represents a request to the near method
type NearRequest struct {
	Profile  string
	GeoPath  GeoPath
	Bearings []Bearing
	Number   int
}

// NearResponse represents a response from the near method
type NearResponse struct {
	Waypoints []NearWaypoint `json:"waypoints"`
}

// NearWaypoint represents a nearest point on a near query
type NearWaypoint struct {
	Location geo.Point `json:"location"`
	Distance float64   `json:"distance"`
	Name     string    `json:"name"`
	Hint     string    `json:"hint"`
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
		service: "nearest",
		geoPath: r.GeoPath,
		options: opts,
	}
}

// URL generates a url for OSRM match request
func (r NearRequest) URL(serverURL string) (string, error) {
	return r.request().URL(serverURL)
}
