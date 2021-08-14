package mock

import (
	"fmt"

	"github.com/shengng325/clean-architecture/v2/db"
	"github.com/shengng325/clean-architecture/v2/dto"
	"github.com/shengng325/clean-architecture/v2/entity"
	"github.com/shengng325/clean-architecture/v2/model"
)

type SpotDB struct {
	spots map[int]*entity.Spot
	maxID int
}

func NewSpotDB() db.SpotDB {
	spotDB := &SpotDB{
		spots: map[int]*entity.Spot{},
		maxID: 0,
	}
	spotDB.seed()
	return spotDB
}

func (u *SpotDB) CreateSpot(spot *dto.Spot) (*entity.Spot, error) {
	newSpot := &entity.Spot{
		ID:          u.maxID + 1,
		DisplayName: spot.DisplayName,
	}
	u.spots[u.maxID+1] = newSpot
	return newSpot, nil
}
func (u *SpotDB) GetSpotByID(id int) (*entity.Spot, error) {
	spot, ok := u.spots[id]
	if !ok {
		return nil, fmt.Errorf("not found")
	}
	return spot, nil
}

func (u *SpotDB) ListSpots() ([]*entity.Spot, error) {
	spots := []*entity.Spot{}
	for _, s := range u.spots {
		spots = append(spots, s)
	}
	return spots, nil
}

func (u *SpotDB) UpdateSpot(spot *entity.Spot) (*entity.Spot, error) {
	if _, ok := u.spots[spot.ID]; !ok {
		return nil, fmt.Errorf("not found")
	}
	u.spots[spot.ID] = spot
	return spot, nil
}
func (u *SpotDB) DeleteSpotByID(id int) bool {
	if _, ok := u.spots[id]; !ok {
		return false
	}
	delete(u.spots, id)
	return true
}

func (s *SpotDB) seed() {
	s.spots[1] = &entity.Spot{
		ID:          1,
		DisplayName: "McD",
		Position: &model.Position{
			Latitude:   10,
			Longtitude: 15,
		},
	}
	s.spots[2] = &entity.Spot{
		ID:          2,
		DisplayName: "KFC",
		Position: &model.Position{
			Latitude:   30,
			Longtitude: 35,
		},
	}
	s.spots[3] = &entity.Spot{
		ID:          3,
		DisplayName: "Texas",
		Position: &model.Position{
			Latitude:   50,
			Longtitude: 55,
		},
	}
	s.maxID = 3
}
