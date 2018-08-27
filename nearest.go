package osrm

import geo "github.com/paulmach/go.geo"

// NearestRequest represents a request to the nearest method
type NearestRequest struct {
	Profile     string
	Coordinates Geometry
	Bearings    []Bearing
	Number      int
}

// NearestResponse represents a response from the nearest method
type NearestResponse struct {
	Waypoints []NearestWaypoint `json:"waypoints"`
}

// NearestWaypoint represents a nearest point on a nearest query
type NearestWaypoint struct {
	Location geo.Point `json:"location"`
	Distance float64   `json:"distance"`
	Name     string    `json:"name"`
	Hint     string    `json:"hint"`
	Nodes    []uint64  `json:"nodes"`
}

type nearestResponseOrError struct {
	responseStatus
	NearestResponse
}

func (r NearestRequest) request() *request {
	opts := options{}
	if r.Number > 0 {
		opts.addInt("number", r.Number)
	}

	if len(r.Bearings) > 0 {
		opts.set("bearings", bearings(r.Bearings))
	}

	return &request{
		profile: r.Profile,
		service: "nearest",
		coords:  r.Coordinates,
		options: opts,
	}
}

// URL generates a url for OSRM match request
func (r NearestRequest) URL(serverURL string) (string, error) {
	return r.request().URL(serverURL)
}
