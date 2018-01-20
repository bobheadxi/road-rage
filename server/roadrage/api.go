package roadrage

import (
	"encoding/json"
	"log"
	"net/http"

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

	lats, ok := req.URL.Query()["lat"]
	if !ok || len(lats) < 1 {
		http.Error(w, "Need latitude", http.StatusBadRequest)
		return
	}
	lat := lats[0]
	lons, ok := req.URL.Query()["lon"]
	if !ok || len(lons) < 1 {
		http.Error(w, "Need longitude", http.StatusBadRequest)
		return
	}
	lon := lons[0]

	seg, err := s.api.GetSegmentAtCoordinate(lat, lon)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	segUnmarshall, _ := json.Marshal(seg)
	w.WriteHeader(http.StatusOK)
	w.Write(segUnmarshall)
}
