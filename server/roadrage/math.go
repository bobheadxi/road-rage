package roadrage

import (
	"strconv"

	"github.com/bobheadxi/road-rage/server/tomtom"
)

func mapCoords(vs []tomtom.Coordinate, f func(tomtom.Coordinate, tomtom.Coordinate) tomtom.Coordinate, center tomtom.Coordinate) []tomtom.Coordinate {
	vsm := make([]tomtom.Coordinate, len(vs))
	for i, v := range vs {
		vsm[i] = f(v, center)
	}
	return vsm
}

func makeRelative(point tomtom.Coordinate, center tomtom.Coordinate) tomtom.Coordinate {
	point.Latitude = (point.Latitude - center.Latitude) * 1000
	point.Longitude = (point.Longitude - center.Longitude) * 1000
	return point
}

func generateGrid(center tomtom.Coordinate, radius float64, interval float64) ([]string, []string) {
	maxLat := center.Latitude + radius
	minLat := center.Latitude - radius
	maxLon := center.Longitude + radius
	minLon := center.Longitude - radius

	var lats []string
	var lons []string
	for lon := minLon; lon < maxLon; lon += interval {
		for lat := minLat; lat < maxLat; lat += interval {
			lats = append(lats, strconv.FormatFloat(lat, 'f', 6, 64))
			lons = append(lons, strconv.FormatFloat(lon, 'f', 6, 64))
		}
	}

	return lats, lons
}
