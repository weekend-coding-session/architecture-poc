package db

import (
	"architecture-poc/dto"
	"architecture-poc/entity"
)

type UserDB interface {
	CreateUser(*dto.User) (*entity.User, error)
	GetUserByID(int) (*entity.User, error)
	UpdateUser(*entity.User) (*entity.User, error)
	DeleteUserByID(id int) bool
}

type SpotDB interface {
	CreateSpot(*dto.Spot) (*entity.Spot, error)
	GetSpotByID(int) (*entity.Spot, error)
	ListSpots() ([]*entity.Spot, error)
	UpdateSpot(*entity.Spot) (*entity.Spot, error)
	DeleteSpotByID(int) bool
}
