package service

import (
	"fmt"

	"github.com/PhillBeck/ingresse-backend/model"
	"github.com/PhillBeck/ingresse-backend/repo"
	"gopkg.in/mgo.v2/bson"
)

// This interface exists only for testing to be mocked for the unit tests
type IUserService interface {
	GetByID(bson.ObjectId) (*model.User, error)
	DeleteByID(bson.ObjectId) error
	Create(*model.User) error
	FindByIdAndReplace(bson.ObjectId, *model.User) error
}

type User struct {
	Repository repo.IUser
}

func NewUserService() *User {
	return &User{
		Repository: repo.NewUserRepo()}
}

func (s *User) GetByID(ID bson.ObjectId) (*model.User, error) {
	return s.Repository.GetByID(ID)
}

func (s *User) DeleteByID(ID bson.ObjectId) error {
	return s.Repository.DeleteByID(ID)
}

func (s *User) Create(user *model.User) error {
	if user.GetID().Hex() == "" {
		user.SetID(bson.NewObjectId())
	}

	return s.Repository.Save(user)
}

func (s *User) FindByIdAndReplace(ID bson.ObjectId, user *model.User) error {
	if user.GetID() != ID {
		return fmt.Errorf("Cannot change property ID")
	}

	_, err := s.Repository.GetByID(ID)
	if err != nil {
		return err
	}

	return s.Repository.Save(user)
}
