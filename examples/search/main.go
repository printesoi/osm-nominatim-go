package main

import (
	"context"
	"fmt"

	nominatim "github.com/printesoi/osm-nominatim-go"
)

func main() {
	ctx := context.Background()
	results, err := nominatim.Search(ctx, "Central Park")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	for _, result := range results {
		fmt.Println("Search Result:", result)
	}
}
