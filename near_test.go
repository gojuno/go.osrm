package osrm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNearRequestOverviewOption(t *testing.T) {
	req := NearRequest{
		Number: 2,
		Bearings: []Bearing{
			{60, 380},
		},
	}
	assert.Equal(
		t,
		"bearings=60%2C380&number=2",
		req.request().options.encode())
}
