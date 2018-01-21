package roadrage

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/bobheadxi/road-rage/server/tomtom"
)

type Server struct {
	api *tomtom.API
}

func New(api *tomtom.API) *Server {
	return &Server{
		api: api,
	}
}

func (s *Server) Run(port string) {
	log.Print("Listening...")
	http.HandleFunc("/build_game", s.buildGame)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func (s *Server) buildGame(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	lat, err := getFloatKey(req.URL.Query(), "lat")
	lon, err := getFloatKey(req.URL.Query(), "lon")
	radius, err := getFloatKey(req.URL.Query(), "radius")
	interval, err := getFloatKey(req.URL.Query(), "interval")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	roadmap, err := s.buildMap(lat, lon, radius, interval)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	unmarshalled, _ := json.Marshal(roadmap)
	w.WriteHeader(http.StatusOK)
	w.Write(unmarshalled)
}

func (s *Server) buildMap(lat float64, lon float64, radius float64, interval float64) (*RoadRageMap, error) {
	center := tomtom.Coordinate{Latitude: lat, Longitude: lon}

	var roads []road
	lats, lons := generateGrid(&center, radius, interval)
	log.Println("Number of points: " + strconv.Itoa(len(lats)))
	for i := 0; i < len(lats); i++ {
		//log.Print(lats[i] + " " + lons[i])
		seg, err := s.api.GetSegmentAtCoordinate(lats[i], lons[i])
		if err != nil {
			continue
		}
		if len(seg.Coordinates.Points) == 0 {
			continue
		}
		roads = append(roads, road{
			Density:     calculateDensity(seg),
			Coordinates: mapCoords(seg.Coordinates.Points, makeRelative, &center),
		})
	}

	roadmap := &RoadRageMap{Center: center, Roads: roads}
	return roadmap, nil
}

func getFloatKey(v url.Values, key string) (float64, error) {
	vals, ok := v[key]
	if !ok || len(vals) < 1 {
		return 0, errors.New("Missing " + key)
	}
	val := vals[0]
	log.Println(key + ": " + val)
	floatVal, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return 0, err
	}

	return floatVal, nil
}
