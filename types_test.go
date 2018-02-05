package osrm

import (
	"encoding/json"
	"testing"

	"github.com/paulmach/go.geo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnmarshalGeoPathFromPointsArray(t *testing.T) {
	gp := GeoPath{}
	jdata := []byte("[[-73.982253,40.742926],[-73.985253,40.742926]]")

	err := json.Unmarshal(jdata, &gp)

	require.Nil(t, err)
	require.Len(t, gp.PointSet, 2)
	require.Equal(t, *geo.NewPointFromLatLng(40.742926, -73.982253), gp.PointSet[0])
	require.Equal(t, *geo.NewPointFromLatLng(40.742926, -73.985253), gp.PointSet[1])
}

func TestUnmarshalGeoPathFromPolyline(t *testing.T) {
	gp := GeoPath{}
	jdata := []byte("\"w{_tlA`a_clCkrDldB~vBcyJ\"")

	err := json.Unmarshal(jdata, &gp)

	require.Nil(t, err)
	require.Len(t, gp.PointSet, 3)
	require.Equal(t, *geo.NewPointFromLatLng(40.71470, -73.990177), gp.PointSet[0])
	require.Equal(t, *geo.NewPointFromLatLng(40.71757, -73.99180), gp.PointSet[1])
	require.Equal(t, *geo.NewPointFromLatLng(40.71565, -73.98575), gp.PointSet[2])
}

func TestPolylineGeoPath(t *testing.T) {
	path := geo.NewPath()
	path.Push(geo.NewPointFromLatLng(40.714701, -73.990177))
	path.Push(geo.NewPointFromLatLng(40.717572, -73.991801))
	path.Push(geo.NewPointFromLatLng(40.715653, -73.985752))
	gp := GeoPath{*path}
	assert.Equal(t, "{aowFrerbM}PbI~Jyd@", gp.Polyline())
}

func TestRequestURLWithEmptyOptions(t *testing.T) {
	req := request{
		profile: "something",
		geoPath: geoPath,
		service: "foobar",
	}
	url, err := req.URL("localhost")
	require.Nil(t, err)
	assert.Equal(t, "localhost/foobar/v1/something/polyline({aowFrerbM}PbI~Jyd@)", url)
}

func TestRequestURLWithOptions(t *testing.T) {
	opts := options{}
	opts.Set("baz", "quux")
	req := request{
		profile: "something",
		geoPath: geoPath,
		service: "foobar",
		options: opts,
	}
	url, err := req.URL("localhost")
	require.Nil(t, err)
	assert.Equal(t, "localhost/foobar/v1/something/polyline({aowFrerbM}PbI~Jyd@)?baz=quux", url)
}

func TestRequestURLWithEmptyService(t *testing.T) {
	req := request{}
	url, err := req.URL("localhost")
	require.NotNil(t, err)
	assert.Equal(t, ErrEmptyServiceName, err)
	assert.Empty(t, url)
}

func TestRequestURLWithEmptyProfile(t *testing.T) {
	req := request{
		service: "foobar",
	}
	url, err := req.URL("localhost")
	require.NotNil(t, err)
	assert.Equal(t, ErrEmptyProfileName, err)
	assert.Empty(t, url)
}

func TestRequestURLWithoutCoords(t *testing.T) {
	req := request{
		profile: "something",
		service: "foobar",
	}
	url, err := req.URL("localhost")
	require.NotNil(t, err)
	assert.Equal(t, ErrNoCoordinates, err)
	assert.Empty(t, url)
}
