package osrm

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/paulmach/go.geo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	geoPath = *NewGeoPathFromPointSet(
		geo.PointSet([]geo.Point{
			{-73.990185, 40.714701},
			{-73.991801, 40.717571},
			{-73.985751, 40.715651},
		}))

	pathRequest = Request{
		Profile: "car",
		GeoPath: geoPath,
	}
)

func fixturedJSON(name string) []byte {
	data, err := ioutil.ReadFile("test_fixtures/" + name + ".json")
	if err != nil {
		log.Fatalf("osrm5: failed to load a fixture %s, err: %s", name, err)
	}
	return data
}

func fixturedHTTPHandler(name string, assertURL func(path, query string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		assertURL(r.URL.Path, r.URL.RawQuery)
		fmt.Fprintln(w, string(fixturedJSON(name)))
	}
}

func TestErrorWithTimeout(t *testing.T) {
	osrm := NewFromURLWithTimeout("http://25.0.0.1", 500*time.Microsecond)

	var nothing Response

	request := Request{
		service: "nothing",
		Profile: "nothing",
		GeoPath: geoPath,
	}

	err := osrm.query(context.Background(), &request, nothing)

	require.NotNil(t, err)
	assert.Equal(t, ErrInternal, err.ErrCode())
}

func TestErrorOnRouteRequest(t *testing.T) {
	ts := httptest.NewServer(fixturedHTTPHandler("route_response_no_route_error", func(path, query string) {
		assert.Equal(t, "/route/v1/car/polyline({aowFrerbM}PbI~Jyd@)", path)
		assert.Equal(t, "annotations=false&continue_straight=true&geometries=polyline6&overview=false&steps=false", query)
	}))
	defer ts.Close()

	osrm := NewFromURL(ts.URL)

	r, err := osrm.Route(context.Background(), RouteRequest{
		Request:     pathRequest,
		Annotations: AnnotationsFalse,
		Steps:       StepsFalse,
		Geometries:  GeometriesPolyline6,
		Overview:    OverviewFalse})

	require.NotNil(t, err)
	assert.Equal(t, ErrNoRoute, err.ErrCode())
	assert.Equal(t, "no route to coordinates", err.Error())
	assert.Nil(t, r)
}

func TestRouteRequest(t *testing.T) {
	ts := httptest.NewServer(fixturedHTTPHandler("route_response_full", func(path, query string) {
		assert.Equal(t, "/route/v1/car/polyline({aowFrerbM}PbI~Jyd@)", path)
		assert.Equal(t, "annotations=true&continue_straight=true&geometries=polyline6&overview=full", query)
	}))
	defer ts.Close()

	osrm := NewFromURL(ts.URL)

	r, err := osrm.Route(context.Background(), RouteRequest{
		Request:     pathRequest,
		Annotations: AnnotationsTrue,
		Geometries:  GeometriesPolyline6,
		Overview:    OverviewFull,
	})

	require := require.New(t)

	require.Nil(err)
	require.NotNil(r)

	// routes
	require.Len(r.Routes, 1)
	route := r.Routes[0]
	require.Equal(float32(1190.5), route.Distance)
	require.Equal(float32(92.2), route.Duration)
	// routes/legs
	require.Len(route.Legs, 2)
	leg0 := route.Legs[0]
	require.Equal(float32(637.5), leg0.Distance)
	require.Equal(float32(58.0), leg0.Duration)
	// routes/annotations
	annotation := leg0.Annotation
	require.Len(annotation.Duration, 14)
	require.Len(annotation.Distance, 14)
	// routes/legs/steps
	require.Len(leg0.Steps, 7)
	// routes/legs/steps[0]
	step0 := leg0.Steps[0]
	require.Equal("driving", step0.Mode)
	require.Equal("", step0.Name)
	require.Equal(float32(5.0), step0.Duration)
	require.Equal(float32(33.1), step0.Distance)
	require.Equal(GeoPath{
		Path: *geo.NewPathFromXYSlice([][]float64{
			{-73.9902, 40.7147},
			{-73.99023, 40.7146},
			{-73.99025, 40.71441},
		}),
	}, step0.Geometry)
}

func TestTableRequest(t *testing.T) {
	ts := httptest.NewServer(fixturedHTTPHandler("table_response_full", func(path, query string) {
		assert.Equal(t, "/table/v1/car/polyline({aowFrerbM}PbI~Jyd@)", path)
		assert.Empty(t, query)
	}))
	defer ts.Close()

	osrm := NewFromURL(ts.URL)

	r, err := osrm.Table(context.Background(), TableRequest{Request: pathRequest})

	require := require.New(t)

	require.Nil(err)
	require.NotNil(r)

	require.Len(r.Durations, 3)
	require.Equal([]float32{0, 39, 46.8}, r.Durations[0])
	require.Equal([]float32{39.5, 0, 34.2}, r.Durations[1])
	require.Equal([]float32{47.2, 34.2, 0}, r.Durations[2])
}

func TestMatchRequest(t *testing.T) {
	ts := httptest.NewServer(fixturedHTTPHandler("match_response_full", func(path, query string) {
		assert.Equal(t, "/match/v1/car/polyline({aowFrerbM}PbI~Jyd@)", path)
		assert.Equal(t, "geometries=polyline6", query)
	}))
	defer ts.Close()

	osrm := NewFromURL(ts.URL)

	r, err := osrm.Match(context.Background(), MatchRequest{
		Request: pathRequest,
	})

	require := require.New(t)

	require.Nil(err)
	require.NotNil(r)

	// matchings
	require.Len(r.Matchings, 1)
	matching := r.Matchings[0]
	require.Equal(0.023898, matching.Confidence)
	require.Equal(float32(1035.3), matching.Distance)
	require.Equal(float32(79.0), matching.Duration)
	// matchings/legs
	require.Len(matching.Legs, 2)
}
