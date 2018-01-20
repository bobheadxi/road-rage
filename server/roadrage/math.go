package roadrage

import "github.com/bobheadxi/road-rage/server/tomtom"

func mapCoords(vs []tomtom.Coordinate, f func(tomtom.Coordinate, tomtom.Coordinate) tomtom.Coordinate, center tomtom.Coordinate) []tomtom.Coordinate {
	vsm := make([]tomtom.Coordinate, len(vs))
	for i, v := range vs {
		vsm[i] = f(v, center)
	}
	return vsm
}

func makeRelative(point tomtom.Coordinate, center tomtom.Coordinate) tomtom.Coordinate {
	point.Latitude = (point.Latitude - center.Latitude) * 10000
	point.Longitude = (point.Longitude - center.Longitude) * 10000
	return point
}
