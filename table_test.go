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
		Sources:      []int{0, 1, 2},
		Destinations: []int{1, 3},
	}
	assert.Equal(t, "destinations=1;3&sources=0;1;2", req.request().options.encode())
}
