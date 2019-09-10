package osrm

import (
	"encoding/json"
	"testing"

	"github.com/paulmach/go.geo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnmarshalGeometryFromPolyline(t *testing.T) {
	gp := Geometry{}
	jdata := []byte("\"w{_tlA`a_clCkrDldB~vBcyJ\"")

	err := json.Unmarshal(jdata, &gp)

	require.Nil(t, err)
	require.Len(t, gp.PointSet, 3)
	require.Equal(t, *geo.NewPointFromLatLng(40.71470, -73.990177), gp.PointSet[0])
	require.Equal(t, *geo.NewPointFromLatLng(40.71757, -73.99180), gp.PointSet[1])
	require.Equal(t, *geo.NewPointFromLatLng(40.71565, -73.98575), gp.PointSet[2])
}

func TestPolylineGeometry(t *testing.T) {
	path := geo.NewPath()
	path.Push(geo.NewPointFromLatLng(40.714701, -73.990177))
	path.Push(geo.NewPointFromLatLng(40.717572, -73.991801))
	path.Push(geo.NewPointFromLatLng(40.715653, -73.985752))
	gp := Geometry{*path}
	assert.Equal(t, "{aowFrerbM}PbI~Jyd@", gp.Polyline())
}

func TestRequestURLWithEmptyOptions(t *testing.T) {
	req := request{
		profile: "something",
		coords:  geometry,
		service: "foobar",
	}
	url, err := req.URL("localhost")
	require.Nil(t, err)
	assert.Equal(t, "localhost/foobar/v1/something/polyline(%7BaowFrerbM%7DPbI~Jyd@)", url)
}

func TestRequestURLWithOptions(t *testing.T) {
	opts := options{}
	opts.set("baz", "quux")
	req := request{
		profile: "something",
		coords:  geometry,
		service: "foobar",
		options: opts,
	}
	url, err := req.URL("localhost")
	require.Nil(t, err)
	assert.Equal(t, "localhost/foobar/v1/something/polyline(%7BaowFrerbM%7DPbI~Jyd@)?baz=quux", url)
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
