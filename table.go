package osrm

// TableRequest represents a request to the table method
type TableRequest struct {
	Request
	Sources, Destinations []int
}

// TableResponse resresents a response from the table method
type TableResponse struct {
	ResponseError
	Durations [][]float32 `json:"durations"`
}

func (r TableRequest) request() *Request {
	r.service = "table"
	r.options = Options{}
	if len(r.Sources) > 0 {
		r.options.AddInt("sources", r.Sources...)
	}
	if len(r.Destinations) > 0 {
		r.options.AddInt("destinations", r.Destinations...)
	}
	return &r.Request
}
