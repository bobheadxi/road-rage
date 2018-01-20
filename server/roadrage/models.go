package roadrage

import "github.com/bobheadxi/road-rage/server/tomtom"

type RoadRageMap struct {
	Roads []road `json:"rooads"`
}

type road struct {
	Density     float32             `json:"density"`
	Coordinates []tomtom.Coordinate `json:"coordinates"`
}
