package service

import (
	"strconv"
	"user-service/db"
)

type UserService struct {
}

type UsersHttpService interface {
	CreateUser(user db.User) error
	GetUsers() []db.User
	FilterUser(dateStart, dateEnd, ageStart, ageEnd string) ([]db.User, error)
}

func (u UserService) CreateUser(user db.User) error {
	err := db.CreateUser(&user)
	if err != nil {
		return err
	}
	return nil
}

func (u UserService) GetUsers() []db.User {
	users, err := db.GetUsers()
	if err != nil {
		return nil
	}
	return users
}

func (u UserService) FilterUser(dateStart, dateEnd, ageStart, ageEnd string) ([]db.User, error) {
	dateStartInt, _ := strconv.ParseInt(dateStart, 10, 64)
	dateEndInt, _ := strconv.ParseInt(dateEnd, 10, 64)
	ageStartInt, _ := strconv.ParseInt(ageStart, 10, 64)
	ageEndInt, _ := strconv.ParseInt(ageEnd, 10, 64)

	defer func() {
		if err := recover(); err != nil {
		}
	}()

	users, err := db.FilterUsers(dateStartInt, dateEndInt, ageStartInt, ageEndInt)
	return users, err
}
