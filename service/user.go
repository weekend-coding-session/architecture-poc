package service

import (
	"fmt"

	"github.com/shengng325/clean-architecture/v2/db"
	"github.com/shengng325/clean-architecture/v2/dto"
	"github.com/shengng325/clean-architecture/v2/entity"
)

type userService struct {
	db db.UserDB
}

func NewUserService(db db.UserDB) User {
	return &userService{
		db: db,
	}
}

func (u *userService) CreateUser(user *dto.User) (*entity.User, error) {
	if user == nil {
		return nil, fmt.Errorf("empty user")
	}
	return u.db.CreateUser(user)
}
func (u *userService) GetUserByID(id int) (*entity.User, error) {
	return u.db.GetUserByID(id)
}
func (u *userService) UpdateUser(user *entity.User) (*entity.User, error) {
	if user == nil {
		return nil, fmt.Errorf("nothing to update")
	}
	return u.db.UpdateUser(user)
}
func (u *userService) DeleteUserByID(id int) bool {
	return u.db.DeleteUserByID(id)
}
