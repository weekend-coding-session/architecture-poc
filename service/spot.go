package service

import (
	"fmt"

	"architecture-poc/db"
	"architecture-poc/dto"
	"architecture-poc/entity"
)

type spotService struct {
	db db.SpotDB
}

func NewSpotService(db db.SpotDB) Spot {
	return &spotService{
		db: db,
	}
}

func (s *spotService) CreateSpot(spot *dto.Spot) (*entity.Spot, error) {
	if spot == nil {
		return nil, fmt.Errorf("empty spot")
	}
	return s.db.CreateSpot(spot)
}

func (s *spotService) GetSpotByID(id int) (*entity.Spot, error) {
	return s.db.GetSpotByID(id)
}

func (s *spotService) ListSpots() ([]*entity.Spot, error) {
	return s.db.ListSpots()
}

func (s *spotService) UpdateSpot(spot *entity.Spot) (*entity.Spot, error) {
	if spot == nil {
		return nil, fmt.Errorf("nothing to update")
	}
	return s.db.UpdateSpot(spot)
}

func (s *spotService) DeleteSpotByID(id int) bool {
	return s.db.DeleteSpotByID(id)
}
