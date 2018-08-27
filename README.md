Go client library for OSRM
==========================

[![GoDoc](https://godoc.org/github.com/gojuno/go.osrm?status.svg)](http://godoc.org/github.com/gojuno/go.osrm)
[![Build Status](https://travis-ci.org/gojuno/go.osrm.svg?branch=master)](https://travis-ci.org/gojuno/go.osrm)
[![Go Report Card](https://goreportcard.com/badge/github.com/gojuno/go.osrm)](https://goreportcard.com/report/github.com/gojuno/go.osrm)

## Description

Currently supported OSRM APIs are:
- [Route service](https://github.com/Project-OSRM/osrm-backend/blob/master/docs/http.md#route-service)
- [Table service](https://github.com/Project-OSRM/osrm-backend/blob/master/docs/http.md#table-service)
- [Match service](https://github.com/Project-OSRM/osrm-backend/blob/master/docs/http.md#match-service)
- [Nearest service](https://github.com/Project-OSRM/osrm-backend/blob/master/docs/http.md#nearest-service)

Not implemeted yet:
- [Trip service](https://github.com/Project-OSRM/osrm-backend/blob/master/docs/http.md#trip-service)
- [Tile service](https://github.com/Project-OSRM/osrm-backend/blob/master/docs/http.md#tile-service)

## Usage

Sample usage:

``` go
package main

import (
	"context"
	"log"
	"time"

	osrm "github.com/gojuno/go.osrm"
	geo "github.com/paulmach/go.geo"
)

func main() {
	client := osrm.NewFromURL("https://router.project-osrm.org")

	ctx, cancelFn := context.WithTimeout(context.Background(), time.Second)
	defer cancelFn()

	resp, err := client.Route(ctx, osrm.RouteRequest{
		Profile: "car",
		Coordinates: osrm.NewGeometryFromPointSet(geo.PointSet{
			{-73.980020, 40.751739},
			{-73.962662, 40.794156},
		}),
		Steps:       osrm.StepsTrue,
		Annotations: osrm.AnnotationsTrue,
		Overview:    osrm.OverviewFalse,
		Geometries:  osrm.GeometriesPolyline6,
	})
	if err != nil {
		log.Fatalf("route failed: %v", err)
	}

	log.Printf("routes are: %+v", resp.Routes)
}
```
