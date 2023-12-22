package osrm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyTripRequestOptions(t *testing.T) {
	req := TripRequest{}
	assert.Equal(
		t,
		"geometries=polyline6",
		req.request().options.encode())
}

func TestTripRequestOptionsWithBearings(t *testing.T) {
	req := TripRequest{
		Bearings: []Bearing{
			{60, 380},
			{45, 180},
		},
	}
	assert.Equal(
		t,
		"bearings=60%2C380%3B45%2C180&geometries=polyline6",
		req.request().options.encode())
}

func TestTripRequestOverviewOption(t *testing.T) {
	req := TripRequest{
		Overview:  OverviewFull,
		Roundtrip: RoundtripFalse,
	}
	assert.Equal(
		t,
		"geometries=polyline6&overview=full&roundtrip=false",
		req.request().options.encode())
}

func TestTripRequestGeometryOption(t *testing.T) {
	req := TripRequest{
		Geometries:  GeometriesPolyline6,
		Annotations: AnnotationsFalse,
		Steps:       StepsFalse,
		Roundtrip:   RoundtripFalse,
	}
	assert.Equal(
		t,
		"annotations=false&geometries=polyline6&roundtrip=false&steps=false",
		req.request().options.encode())
}

func TestTripRequestDestinationOption(t *testing.T) {
	req := TripRequest{
		Overview:    OverviewFull,
		Destination: DestinationLast,
	}
	assert.Equal(
		t,
		"destination=last&geometries=polyline6&overview=full",
		req.request().options.encode())
}
