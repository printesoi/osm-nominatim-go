package main

import (
	"context"
	"fmt"

	nominatim "github.com/printesoi/osm-nominatim-go"
)

func main() {
	ctx := context.Background()
	result, err := nominatim.DetailsWithPlaceID(ctx, 240109189) // PlaceID for the Empire State Building
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Details Result:", result)
}
