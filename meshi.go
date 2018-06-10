package meshi

import (
	"log"

	"golang.org/x/net/context"
	"googlemaps.github.io/maps"
)

const FREE_QUERY_PER_SECOND = 2

func main() {
}

// meshi searches some restraunt with given name and returns
func Meshi(api_key string, lat float64, lng float64, rad uint, keyword string) maps.PlacesSearchResponse {
	c, err := maps.NewClient(maps.WithAPIKey(api_key), maps.WithRateLimit(FREE_QUERY_PER_SECOND))
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}

	r := &maps.NearbySearchRequest{
		Location: &maps.LatLng{
			Lat: lat,
			Lng: lng,
		},
		Radius:  rad,
		Keyword: keyword,
		Type:    maps.PlaceTypeRestaurant,
		OpenNow: false,
	}

	response, err := c.NearbySearch(context.Background(), r)
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}
	return response
}
