package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"architecture-poc/common"
	"architecture-poc/model"
	"architecture-poc/service"
)

type Random struct {
	randomService service.Random
}

func NewRandomController(rs service.Random) *Random {
	return &Random{
		randomService: rs,
	}
}

func (c *Random) GetRandomSpots(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	lat, errLat := strconv.Atoi(vars["lat"])
	long, errLong := strconv.Atoi(vars["long"])
	if errLat != nil || errLong != nil {
		http.Error(w, "Invalid pos", http.StatusNotFound)
		return
	}
	pos := &model.Position{
		Latitude:   lat,
		Longtitude: long,
	}
	spot, err := c.randomService.GetRandomSpot(pos)
	if err != nil {
		http.Error(w, "Cant get spot", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(spot)
	common.Log("Random spot found!")
}

func (c *Random) GetNearestSpot(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	lat, errLat := strconv.Atoi(vars["lat"])
	long, errLong := strconv.Atoi(vars["long"])
	if errLat != nil || errLong != nil {
		http.Error(w, "Invalid pos", http.StatusNotFound)
		return
	}
	pos := &model.Position{
		Latitude:   lat,
		Longtitude: long,
	}
	fmt.Println(pos)
	spot, err := c.randomService.GetNearestSpot(pos)
	if err != nil {
		http.Error(w, "Cant get spot", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(spot)
	common.Log("Nearest spot found!")
}
