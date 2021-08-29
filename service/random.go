package service

import (
	"fmt"
	"math/rand"

	"architecture-poc/common"
	"architecture-poc/entity"
	"architecture-poc/model"
)

type randomService struct {
	userService User
	spotService Spot
}

func NewRandomService(us User, ss Spot) Random {
	return &randomService{
		userService: us,
		spotService: ss,
	}
}

func (r *randomService) GetNearestSpot(pos *model.Position) (*entity.Spot, error) {
	if pos == nil {
		return nil, fmt.Errorf("no position")
	}
	spotsAvailable, err := r.spotService.ListSpots()
	if err != nil || len(spotsAvailable) == 0 {
		return nil, fmt.Errorf("no spots found")
	}
	var nearest *entity.Spot
	for i, s := range spotsAvailable {
		if i == 0 {
			nearest = s
			continue
		}
		if r.getDistance(pos, s.Position) < r.getDistance(pos, nearest.Position) {
			nearest = s
		}
	}
	return nearest, nil
}
func (r *randomService) GetRandomSpot(pos *model.Position) (*entity.Spot, error) {
	if pos == nil {
		return nil, fmt.Errorf("no position")
	}
	spotsAvailable, err := r.spotService.ListSpots()
	if err != nil || len(spotsAvailable) == 0 {
		return nil, fmt.Errorf("no spots found")
	}
	randIdx := rand.Intn(len(spotsAvailable))
	return spotsAvailable[randIdx], nil
}

func (r *randomService) getDistance(from *model.Position, to *model.Position) int {
	latDist := common.Abs(to.Latitude - from.Latitude)
	longDist := common.Abs(to.Longtitude - from.Longtitude)
	return int(latDist + longDist)
}
