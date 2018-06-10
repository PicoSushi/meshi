package main

import (
	"log"
	"os"

	"github.com/kr/pretty"
	"golang.org/x/net/context"
	"googlemaps.github.io/maps"
)

const FREE_QUERY_PER_SECOND = 45

func main() {
	api_key := os.Getenv("GOOGLE_MAPS_API_KEY")

	c, err := maps.NewClient(maps.WithAPIKey(api_key), maps.WithRateLimit(FREE_QUERY_PER_SECOND))
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}

	shinjuku_station := &maps.LatLng{
		Lat: 35.690921,
		Lng: 139.70025799999996,
	}

	r := &maps.NearbySearchRequest{
		Location: shinjuku_station,
		Radius:   300,
		Keyword:  "いちごパフェ",
		OpenNow:  false,
	}

	response, err := c.NearbySearch(context.Background(), r)
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}

	for _, result := range response.Results {
		pretty.Println(result.Name)
	}
}
