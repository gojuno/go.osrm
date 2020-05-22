package osrm

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	geo "github.com/paulmach/go.geo"
	geojson "github.com/paulmach/go.geojson"
)

const (
	polyline5Factor = 1.0e5
	polyline6Factor = 1.0e6
)

// Geometry represents a points set
type Geometry struct {
	geo.Path
}

// NewGeometryFromPath creates a geometry from a path.
func NewGeometryFromPath(path geo.Path) Geometry {
	return Geometry{path}
}

// NewGeometryFromPointSet creates a geometry from points set.
func NewGeometryFromPointSet(ps geo.PointSet) Geometry {
	return NewGeometryFromPath(geo.Path{PointSet: ps})
}

// Polyline generates a polyline in Google format
func (g *Geometry) Polyline(factor ...int) string {
	if len(factor) == 0 {
		return g.Encode(polyline5Factor)
	}

	return g.Encode(factor[0])
}

// UnmarshalJSON parses a geo path from points set or a polyline
func (g *Geometry) UnmarshalJSON(b []byte) error {
	if len(b) == 0 {
		return nil
	}

	var encoded string
	if err := json.Unmarshal(b, &encoded); err == nil {
		g.Path = *geo.NewPathFromEncoding(encoded, polyline6Factor)
		return nil
	}

	geom, err := geojson.UnmarshalGeometry(b)
	if err != nil {
		return fmt.Errorf("failed to unmarshal geojson geometry, err: %v", err)
	}
	if !geom.IsLineString() {
		return fmt.Errorf("unexpected geometry type: %v", geom.Type)
	}
	g.Path = *geo.NewPathFromXYSlice(geom.LineString)

	return nil
}

// MarshalJSON generates a polyline in Google polyline6 format
func (g Geometry) MarshalJSON() ([]byte, error) {
	return json.Marshal(g.Polyline(polyline6Factor))
}

// Tidy represents a tidy param for osrm5 match request
type Tidy string

// Supported tidy param values
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

// Supported annotations param values
const (
	AnnotationsTrue        Annotations = "true"
	AnnotationsFalse       Annotations = "false"
	AnnotationsNodes       Annotations = "nodes"
	AnnotationsDistance    Annotations = "distance"
	AnnotationsDuration    Annotations = "duration"
	AnnotationsDatasources Annotations = "datasources"
	AnnotationsWeight      Annotations = "weight"
	AnnotationsSpeed       Annotations = "speed"
)

// String returns Annotations as a string
func (a Annotations) String() string {
	return string(a)
}

// Steps represents a steps param for osrm5 request
type Steps string

// Supported steps param values
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

// Supported gaps param values
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

// Supported geometries param values
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

// ContinueStraight represents continue_straight OSRM routing parameter
type ContinueStraight string

// ContinueStraight values
const (
	ContinueStraightDefault ContinueStraight = "default"
	ContinueStraightTrue    ContinueStraight = "true"
	ContinueStraightFalse   ContinueStraight = "false"
)

// String returns ContinueStraight as string
func (c ContinueStraight) String() string {
	return string(c)
}

// request contains parameters for OSRM query
type request struct {
	profile string
	coords  Geometry
	service string
	options options
}

// URLPath generates a url path for OSRM request
func (r *request) URLPath() (string, error) {
	if r.service == "" {
		return "", ErrEmptyServiceName
	}
	if r.profile == "" {
		return "", ErrEmptyProfileName
	}
	if r.coords.Length() == 0 {
		return "", ErrNoCoordinates
	}
	// {service}/{version}/{profile}/{coordinates}[.{format}]?option=value&option=value
	u := strings.Join([]string{
		r.service, // service
		version,   // version
		r.profile, // profile
		"polyline(" + url.PathEscape(r.coords.Polyline(polyline5Factor)) + ")", // coordinates
	}, "/")
	if len(r.options) > 0 {
		u += "?" + r.options.encode() // options
	}
	return u, nil
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
