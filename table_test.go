package osrm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTableRequestOptions(t *testing.T) {
	cases := []struct {
		name        string
		request     TableRequest
		expectedURI string
	}{
		{
			name:        "empty",
			request:     TableRequest{},
			expectedURI: "",
		},
		{
			name: "with sources and destinations",
			request: TableRequest{
				Sources:      []int{0, 1, 2},
				Destinations: []int{1, 3},
			},
			expectedURI: "destinations=1;3&sources=0;1;2",
		},
		{
			name: "scale_factor",
			request: TableRequest{
				ScaleFactor: 0.8,
				GeneralOptions: GeneralOptions{
					Exclude: []string{"toll"},
				},
			},
			expectedURI: "exclude=toll&scale_factor=0.8",
		},
		{
			name: "fallback_coordinate",
			request: TableRequest{
				FallbackSpeed: 11.5,
				GeneralOptions: GeneralOptions{
					Hints: []string{"a", "b"},
				},
			},
			expectedURI: "fallback_speed=11.5&hints=a;b",
		},
		{
			name: "fallback_coordinate",
			request: TableRequest{
				FallbackSpeed:      11.5,
				FallbackCoordinate: FallbackCoordinateInput,
			},
			expectedURI: "fallback_coordinate=input&fallback_speed=11.5",
		},
		{
			name: "annotations",
			request: TableRequest{
				Annotations:   AnnotationsDurationDistance,
				FallbackSpeed: 11.5,
			},
			expectedURI: "annotations=duration%2Cdistance&fallback_speed=11.5",
		},
		{
			name: "fallback_coordinate snapped",
			request: TableRequest{
				FallbackSpeed:      11.5,
				FallbackCoordinate: FallbackCoordinateSnapped,
			},
			expectedURI: "fallback_coordinate=snapped&fallback_speed=11.5",
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.expectedURI, c.request.request().options.encode())
		})
	}
}
