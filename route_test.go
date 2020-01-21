package osrm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyRouteRequestOptions(t *testing.T) {
	req := RouteRequest{}
	assert.Equal(
		t,
		"geometries=polyline6",
		req.request().options.encode())
}

func TestRouteRequestOptionsWithBearings(t *testing.T) {
	req := RouteRequest{
		GeneralOptions: GeneralOptions{
			Bearings: []Bearing{
				{60, 380},
				{45, 180},
			},
		},
		ContinueStraight: ContinueStraightTrue,
	}
	assert.Equal(
		t,
		"bearings=60%2C380%3B45%2C180&continue_straight=true&geometries=polyline6",
		req.request().options.encode())
}

func TestRouteRequestOverviewOption(t *testing.T) {
	req := RouteRequest{
		Overview:         OverviewFull,
		ContinueStraight: ContinueStraightTrue,
	}
	assert.Equal(
		t,
		"continue_straight=true&geometries=polyline6&overview=full",
		req.request().options.encode())
}

func TestRouteRequestGeometryOption(t *testing.T) {
	req := RouteRequest{
		Geometries:       GeometriesPolyline6,
		Annotations:      AnnotationsFalse,
		Steps:            StepsFalse,
		ContinueStraight: ContinueStraightTrue,
	}
	assert.Equal(
		t,
		"annotations=false&continue_straight=true&geometries=polyline6&steps=false",
		req.request().options.encode())
}
