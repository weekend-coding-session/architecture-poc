package service

import (
	"github.com/shengng325/clean-architecture/v2/dto"
	"github.com/shengng325/clean-architecture/v2/entity"
	"github.com/shengng325/clean-architecture/v2/model"
)

type User interface {
	CreateUser(*dto.User) (*entity.User, error)
	GetUserByID(int) (*entity.User, error)
	UpdateUser(*entity.User) (*entity.User, error)
	DeleteUserByID(id int) bool
}

type Spot interface {
	CreateSpot(*dto.Spot) (*entity.Spot, error)
	GetSpotByID(int) (*entity.Spot, error)
	ListSpots() ([]*entity.Spot, error)
	UpdateSpot(*entity.Spot) (*entity.Spot, error)
	DeleteSpotByID(int) bool
}

type Random interface {
	GetNearestSpot(*model.Position) (*entity.Spot, error)
	GetRandomSpot(*model.Position) (*entity.Spot, error)
}
