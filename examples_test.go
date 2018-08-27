package osrm_test

import (
	"context"
	"fmt"
	"log"

	osrm "github.com/gojuno/go.osrm"
	geo "github.com/paulmach/go.geo"
)

func ExampleOSRM_Route() {
	client := osrm.NewFromURL("https://router.project-osrm.org")

	resp, err := client.Route(context.Background(), osrm.RouteRequest{
		Profile: "car",
		Coordinates: osrm.NewGeometryFromPointSet(geo.PointSet{
			{-73.87946, 40.75833},
			{-73.87925, 40.75837},
			{-73.87918, 40.75837},
			{-73.87911, 40.75838},
		}),
		Steps:       osrm.StepsTrue,
		Annotations: osrm.AnnotationsTrue,
		Overview:    osrm.OverviewFalse,
		Geometries:  osrm.GeometriesPolyline6,
	})
	if err != nil {
		log.Fatalf("route failed: %v", err)
	}

	fmt.Println(len(resp.Routes))

	// Output:
	// 1
}
