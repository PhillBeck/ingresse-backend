package model

import (
	"time"

	"github.com/PhillBeck/golang-odm"
)

type User struct {
	odm.DocumentBase `bson:",inline"`
	Name             string    `json:"name" bson:"name" binding:"Required"`
	BirthDate        time.Time `json:"birthDate" bson:"birthDate"`
	CPF              string    `json:"cpf" bson:"cpf" binding:"Required"`
	Username         string    `json:"username" bson:"username" binding:"Required"`
}
