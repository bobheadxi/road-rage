package roadrage

import "github.com/bobheadxi/road-rage/server/tomtom"

type RoadRageMap struct {
	Center tomtom.Coordinate `json:"center"`
	Roads  []road            `json:"roads"`
}

type road struct {
	Density     float64             `json:"density"`
	Coordinates []tomtom.Coordinate `json:"coordinates"`
}
