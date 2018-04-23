package repo

import (
	"github.com/PhillBeck/golang-odm"
	"github.com/PhillBeck/ingresse-backend/conf"
	"gopkg.in/mgo.v2/bson"
)

var databaseName string

func init() {
	odm.MongodbURI = conf.GetMongoURI()
	databaseName = conf.GetMongoDatabaseName()
}

// IODMRepository is an interface wrapping odm.Repository.
// This interface is created for mocking purposes.
type IODMRepository interface {
	Save(odm.IEntity) error
	Delete(odm.IEntity) error
	GetByID(bson.ObjectId, odm.IEntity) error
}
