package osrm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyTripRequestOptions(t *testing.T) {
	req := TripRequest{}
	assert.Equal(
		t,
		"destination=any&geometries=polyline6&roundtrip=true&source=any",
		req.request().options.encode())
}

func TestTripRequestOptionsWithRoundtrip(t *testing.T) {
	req := TripRequest{
		Roundtrip: RoundtripFalse,
	}
	assert.Equal(
		t,
		"destination=any&geometries=polyline6&roundtrip=false&source=any",
		req.request().options.encode())
}

func TestTripRequestOptionsWithSource(t *testing.T) {
	req := TripRequest{
		Source: SourceFirst,
	}
	assert.Equal(
		t,
		"destination=any&geometries=polyline6&roundtrip=true&source=first",
		req.request().options.encode())
}

func TestTripRequestOptionsWithDestination(t *testing.T) {
	req := TripRequest{
		Destination: DestinationLast,
	}
	assert.Equal(
		t,
		"destination=last&geometries=polyline6&roundtrip=true&source=any",
		req.request().options.encode())
}

func TestTripRequestOptions(t *testing.T) {
	req := TripRequest{
		Roundtrip:   RoundtripTrue,
		Destination: DestinationLast,
	}
	assert.Equal(
		t,
		"destination=last&geometries=polyline6&roundtrip=true&source=any",
		req.request().options.encode())
}

func TestUnsupportedTripRequestOptionsA(t *testing.T) {
	req := TripRequest{
		Roundtrip:   RoundtripFalse,
		Source:      SourceFirst,
		Destination: DestinationAny,
	}
	assert.Equal(
		t,
		false,
		req.IsSupported())
}

func TestUnsupportedTripRequestOptionsB(t *testing.T) {
	req := TripRequest{
		Roundtrip:   RoundtripFalse,
		Source:      SourceAny,
		Destination: DestinationLast,
	}
	assert.Equal(
		t,
		false,
		req.IsSupported())
}

func TestUnsupportedTripRequestOptionsC(t *testing.T) {
	req := TripRequest{
		Roundtrip:   RoundtripFalse,
		Source:      SourceAny,
		Destination: DestinationAny,
	}
	assert.Equal(
		t,
		false,
		req.IsSupported())
}
