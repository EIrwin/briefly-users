package services

import (
	"github.com/eirwin/briefly-users/data"
	"log"
)

func CreateUser(req *CreateUserRequest) (data.User, error) {
	user := data.User{HourlyRate: req.HourlyRate, Salary: req.Salary}
	err := data.CreateUser(&user)
	if err != nil {
		log.Fatal(err)
	}
	return user, nil
}

func GetUser(req *GetUserRequest) (data.User, error) {
	var user data.User
	user, err := data.GetUser(req.Id)
	if err != nil {
		log.Fatal(err)
	}
	return user, nil
}

type CreateUserRequest struct {
	HourlyRate  float32
	Salary      float32
	IntervalAmt float32
}

type GetUserRequest struct {
	Id string
}
