package osrm

import (
	"encoding/json"
	"testing"

	"github.com/paulmach/go.geo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnmarshalGeometryFromGeojson(t *testing.T) {
	var g Geometry
	in := []byte(`{"type": "LineString", "coordinates": [[-73.982253,40.742926],[-73.985253,40.742926]]}`)

	err := json.Unmarshal(in, &g)

	require.Nil(t, err)
	require.Len(t, g.PointSet, 2)
	require.Equal(t, *geo.NewPoint(-73.982253, 40.742926), g.PointSet[0])
	require.Equal(t, *geo.NewPoint(-73.985253, 40.742926), g.PointSet[1])
}

func TestUnmarshalGeometryFromPolyline(t *testing.T) {
	var g Geometry
	in := []byte(`"nvnalCui}okAkgpk@u}hQf}_l@mbpL"`)

	err := json.Unmarshal(in, &g)

	require.Nil(t, err)
	require.Len(t, g.PointSet, 3)
	require.Equal(t, *geo.NewPoint(40.123563, -73.965432), g.PointSet[0])
	require.Equal(t, *geo.NewPoint(40.423574, -73.235698), g.PointSet[1])
	require.Equal(t, *geo.NewPoint(40.645325, -73.973462), g.PointSet[2])
}

func TestUnmarshalGeometryFromNull(t *testing.T) {
	var g Geometry
	in := []byte(`null`)
	err := json.Unmarshal(in, &g)

	require.Nil(t, err)
	require.Equal(t, 0, len(g.Path.PointSet))
}

func TestUnmarshalGeometryFromEmptyJSON(t *testing.T) {
	var g Geometry
	in := []byte(`{}`)
	err := json.Unmarshal(in, &g)

	require.Error(t, err)
}

func TestPolylineGeometry(t *testing.T) {
	g := Geometry{
		Path: geo.Path{
			PointSet: []geo.Point{
				{40.123563, -73.965432},
				{40.423574, -73.235698},
				{40.645325, -73.973462},
			},
		},
	}

	bytes, err := json.Marshal(g)
	require.NoError(t, err)

	assert.Equal(t, `"nvnalCui}okAkgpk@u}hQf}_l@mbpL"`, string(bytes))
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
