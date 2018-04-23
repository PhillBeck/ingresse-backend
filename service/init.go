package service

import "gopkg.in/mgo.v2/bson"

// PaginationOptions defines the parameters to perform pagination on a collection of documents
type PaginationOptions struct {
	Query          bson.M
	Page           int
	RecordsPerPage int
}
