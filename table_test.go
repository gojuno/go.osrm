package osrm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyTableRequestOptions(t *testing.T) {
	req := TableRequest{}
	assert.Empty(t, req.request().options.encode())
}

func TestNotEmptyTableRequestOptions(t *testing.T) {
	req := TableRequest{
		Sources:            []int{0, 1, 2},
		Destinations:       []int{1, 3},
		Annotations:        AnnotationsDuration,
		FallbackSpeed:      45,
		FallbackCoordinate: FallbackCoordinateSnapped,
		ScaleFactor:        1.052,
	}
	assert.Equal(t, "annotations=duration&destinations=1;3&fallback_coordinate=snapped&fallback_speed=45&scale_factor=1.052&sources=0;1;2", req.request().options.encode())
}
