package lib

import (
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

func ParseObjectID(sID string) (bson.ObjectId, error) {
	if !bson.IsObjectIdHex(sID) {
		return "", fmt.Errorf("invalid id")
	}

	return bson.ObjectIdHex(sID), nil
}
