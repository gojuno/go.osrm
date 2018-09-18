package osrm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyMatchRequestOptions(t *testing.T) {
	cases := []struct {
		name        string
		request     MatchRequest
		expectedURI string
	}{
		{
			name:        "empty",
			expectedURI: "geometries=polyline6",
		},
		{
			name: "with timestamps and radiuses",
			request: MatchRequest{
				Timestamps: []int64{0, 1, 2},
				Radiuses:   []float64{0.123123, 0.12312},
			},
			expectedURI: "geometries=polyline6&radiuses=0.123123;0.12312&timestamps=0;1;2",
		},
		{
			name: "with gaps and tidy",
			request: MatchRequest{
				Timestamps: []int64{0, 1, 2},
				Radiuses:   []float64{0.123123, 0.12312},
				Gaps:       GapsSplit,
				Tidy:       TidyTrue,
			},
			expectedURI: "gaps=split&geometries=polyline6&radiuses=0.123123;0.12312&tidy=true&timestamps=0;1;2",
		},
		{
			name: "with hints",
			request: MatchRequest{
				Hints: []string{"a", "b", "c", "d"},
			},
			expectedURI: "geometries=polyline6&hints=a;b;c;d",
		},
		{
			name: "with bearings",
			request: MatchRequest{
				Bearings: []Bearing{
					{0, 20}, {10, 20},
				},
			},
			expectedURI: "bearings=0%2C20%3B10%2C20&geometries=polyline6",
		},
		{
			name: "custom overview option",
			request: MatchRequest{
				Overview:    OverviewSimplified,
				Geometries:  GeometriesGeojson,
				Annotations: AnnotationsFalse,
				Tidy:        TidyFalse,
				Steps:       StepsFalse,
			},
			expectedURI: "annotations=false&geometries=geojson&overview=simplified&steps=false&tidy=false",
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.expectedURI, c.request.request().options.encode())
		})
	}
}
