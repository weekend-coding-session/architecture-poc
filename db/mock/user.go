package mock

import (
	"fmt"

	"github.com/shengng325/clean-architecture/v2/db"
	"github.com/shengng325/clean-architecture/v2/dto"
	"github.com/shengng325/clean-architecture/v2/entity"
)

type UserDB struct {
	users map[int]*entity.User
	maxID int
}

func NewUserDB() db.UserDB {
	userDB := &UserDB{
		users: map[int]*entity.User{},
		maxID: 0,
	}
	userDB.seed()
	return userDB
}

func (u *UserDB) CreateUser(user *dto.User) (*entity.User, error) {
	newUser := &entity.User{
		ID:          u.maxID + 1,
		DisplayName: user.DisplayName,
	}
	u.users[u.maxID+1] = newUser
	return newUser, nil
}
func (u *UserDB) GetUserByID(id int) (*entity.User, error) {
	user, ok := u.users[id]
	if !ok {
		return nil, fmt.Errorf("not found")
	}
	return user, nil
}
func (u *UserDB) UpdateUser(user *entity.User) (*entity.User, error) {
	if _, ok := u.users[user.ID]; !ok {
		return nil, fmt.Errorf("not found")
	}
	u.users[user.ID] = user
	return user, nil
}
func (u *UserDB) DeleteUserByID(id int) bool {
	if _, ok := u.users[id]; !ok {
		return false
	}
	delete(u.users, id)
	return true
}

func (u *UserDB) seed() {
	u.users[1] = &entity.User{
		ID:          1,
		DisplayName: "Bob-userOne",
	}
	u.users[2] = &entity.User{
		ID:          2,
		DisplayName: "John-userTwo",
	}
	u.users[3] = &entity.User{
		ID:          3,
		DisplayName: "Xin-userThree",
	}
	u.maxID = 3
}
