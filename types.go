package osrm

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/paulmach/go.geo"
)

const polyline6Factor = 1.0e6

// GeoPath represents a points set
type GeoPath struct {
	geo.Path
}

// NewGeoPathFromPointSet create a geo path from points set
func NewGeoPathFromPointSet(s geo.PointSet) *GeoPath {
	return &GeoPath{
		Path: geo.Path{
			PointSet: s,
		},
	}
}

// Tidy represents a tidy param for osrm5 match request
type Tidy string

const (
	TidyTrue  Tidy = "true"
	TidyFalse Tidy = "false"
)

// String returns Tidy as a string
func (t Tidy) String() string {
	return string(t)
}

// Annotations represents a annotations param for osrm5 request
type Annotations string

const (
	AnnotationsTrue  Annotations = "true"
	AnnotationsFalse Annotations = "false"
)

// String returns Annotations as a string
func (a Annotations) String() string {
	return string(a)
}

// Steps represents a steps param for osrm5 request
type Steps string

const (
	StepsTrue  Steps = "true"
	StepsFalse Steps = "false"
)

// String returns Steps as a string
func (s Steps) String() string {
	return string(s)
}

// Gaps represents a gaps param for osrm5 match request
type Gaps string

const (
	GapsSplit  Gaps = "split"
	GapsIgnore Gaps = "ignore"
)

// String returns Gaps as a string
func (g Gaps) String() string {
	return string(g)
}

// Geometries represents a geometries param for osrm5
type Geometries string

const (
	GeometriesPolyline6 Geometries = "polyline6"
	GeometriesGeojson   Geometries = "geojson"
)

// String returns Geometries as a string
func (g Geometries) String() string {
	return string(g)
}

// Overview represents level of overview of geometry in a response
type Overview string

// Available overview values
const (
	OverviewSimplified Overview = "simplified"
	OverviewFull       Overview = "full"
	OverviewFalse      Overview = "false"
)

// String returns Overview as a string
func (o Overview) String() string {
	return string(o)
}

// Request contains parameters for OSRM query
type Request struct {
	Profile string
	GeoPath GeoPath

	service string
	options Options
}

// Response contains properties from OSRM query
type Response interface {
	Error
}

// URL generates a url for OSRM request
func (r *Request) URL(serverURL string) (string, error) {
	if r.service == "" {
		return "", ErrEmptyServiceName
	}
	if r.Profile == "" {
		return "", ErrEmptyProfileName
	}
	if r.GeoPath.Length() == 0 {
		return "", ErrNoCoordinates
	}
	// http://{server}/{service}/{version}/{profile}/{coordinates}[.{format}]?option=value&option=value
	url := strings.Join([]string{
		serverURL, // server
		r.service, // service
		version,   // version
		r.Profile, // profile
		"polyline(" + r.GeoPath.Polyline() + ")", // coordinates
	}, "/")
	if len(r.options) > 0 {
		url += "?" + r.options.Encode() // options
	}
	return url, nil
}

// Polyline generates a polyline in Google format
// It uses default factor because of OSRM5 doesn't support polyline6 as coordinates
func (g *GeoPath) Polyline() string {
	return g.Encode()
}

// UnmarshalJSON parses a geo path from points set or a polyline
func (g *GeoPath) UnmarshalJSON(b []byte) (err error) {
	var encoded string
	if err = json.Unmarshal(b, &encoded); err == nil {
		g.Path = *geo.NewPathFromEncoding(encoded, polyline6Factor)
		return
	}
	return json.Unmarshal(b, &g.PointSet)
}

// Bearing limits the search to segments with given bearing in degrees towards true north in clockwise direction.
type Bearing struct {
	Value, Range uint16
}

func (b Bearing) String() string {
	return fmt.Sprintf("%d,%d", b.Value, b.Range)
}

func bearings(br []Bearing) string {
	s := make([]string, len(br))
	for i, b := range br {
		s[i] = b.String()
	}
	return strings.Join(s, ";")
}
