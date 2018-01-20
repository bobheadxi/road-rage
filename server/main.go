package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/bobheadxi/road-rage/server/roadrage"
	"github.com/bobheadxi/road-rage/server/tomtom"
)

/**
Vancouver
Latitude‎: ‎49.246292
Longitude‎: ‎-123.116226
**/

func main() {
	api := tomtom.New()
	seg, err := api.GetSegmentAtCoordinate("49.246292", "-123.116226")
	if err != nil {
		log.Fatal("Fuck" + err.Error())
	}
	segUnmarshall, _ := json.Marshal(seg)
	fmt.Println(string(segUnmarshall))

	server := roadrage.New(api)
	server.Run("8000")
}
