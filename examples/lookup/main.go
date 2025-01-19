package main

import (
	"context"
	"fmt"

	nominatim "github.com/printesoi/osm-nominatim-go"
)

func main() {
	ctx := context.Background()
	result, err := nominatim.Lookup(ctx, "way", 123456789) // Example OSM type and ID
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Lookup Result:", result)
}
