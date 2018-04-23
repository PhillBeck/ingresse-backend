package repo

import (
	"github.com/PhillBeck/golang-odm"
	"github.com/PhillBeck/ingresse-backend/model"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	CollectionName string
	Repo           *odm.Repository
}

type IUser interface {
	GetByID(bson.ObjectId) (*model.User, error)
	Save(*model.User) error
	DeleteByID(bson.ObjectId) error
	Paginate(bson.M, int, int) ([]*model.User, *odm.PaginationInfo, error)
}

func NewUserRepo() *User {
	col := "users"
	return &User{
		CollectionName: col,
		Repo:           odm.NewRepository(databaseName, col)}
}

func (u *User) GetByID(ID bson.ObjectId) (*model.User, error) {
	user := model.User{}
	err := u.Repo.GetByID(ID, &user)

	return &user, err
}

func (u *User) Save(user *model.User) error {
	return u.Repo.Save(user)
}

func (u *User) DeleteByID(ID bson.ObjectId) error {
	user, err := u.GetByID(ID)
	if err != nil {
		return err
	}

	return u.Repo.Delete(user)
}

func (u *User) Paginate(query bson.M, perPage, page int) ([]*model.User, *odm.PaginationInfo, error) {
	resultSet, info, err := u.Repo.Paginate(query, perPage, page)
	if err != nil {
		return nil, nil, err
	}

	var documents []*model.User

	err = resultSet.All(&documents)

	return documents, info, err
}
