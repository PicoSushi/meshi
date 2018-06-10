package main

import (
	"fmt"
	"os"

	"github.com/picosushi/meshi"
)

func run() int {
	api_key := os.Getenv("GOOGLE_MAPS_API_KEY")
	response := meshi.Meshi(api_key, 35.690921, 139.700258, 500, "肉")

	for _, result := range response.Results {
		fmt.Println(result.Name)
	}
	if response.NextPageToken != "" {
		fmt.Println(response.NextPageToken)
	}
	return 0
}

func main() {
	os.Exit(run())
}
