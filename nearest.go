package osrm

import geo "github.com/paulmach/go.geo"

// NearestRequest represents a request to the nearest method
type NearestRequest struct {
	GeneralOptions
	Profile     string
	Coordinates Geometry
	Number      int
}

// NearestResponse represents a response from the nearest method
type NearestResponse struct {
	ResponseStatus
	Waypoints []NearestWaypoint `json:"waypoints"`
}

// NearestWaypoint represents a nearest point on a nearest query
type NearestWaypoint struct {
	Waypoint
	Nodes []uint64 `json:"nodes"`
}

func (r NearestRequest) request() *request {
	opts := options{}
	if r.Number > 0 {
		opts.addInt("number", r.Number)
	}

	opts = r.GeneralOptions.options(opts)

	return &request{
		profile: r.Profile,
		service: "nearest",
		coords:  r.Coordinates,
		options: opts,
	}
}

type Waypoint struct {
	Location geo.Point `json:"location"`
	Distance float64   `json:"distance"`
	Name     string    `json:"name"`
	Hint     string    `json:"hint"`
}
