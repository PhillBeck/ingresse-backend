package query

import (
	"fmt"
	"strings"

	"github.com/PhillBeck/ingresse-backend/service"
	"gopkg.in/mgo.v2/bson"
)

func ParseQuery(search string, page, perPage int) (service.PaginationOptions, error) {
	options := service.PaginationOptions{
		Page:           1,
		RecordsPerPage: 10}

	if page > 0 {
		options.Page = page
	}

	if perPage > 0 {
		options.RecordsPerPage = perPage
	}

	if search != "" {
		query, err := parseSearch(search)
		if err != nil {
			return options, err
		}

		options.Query = query
	}

	return options, nil
}

func parseSearch(query string) (bson.M, error) {
	searchOptions := bson.M{}
	terms := strings.Split(query, ";")

	for _, term := range terms {
		parts := strings.Split(term, ":")
		if len(parts) != 2 {
			err := fmt.Errorf("Invalid Query")
			return bson.M{}, err
		}

		searchOptions[parts[0]] = bson.M{"$regex": parts[1]}
	}

	return searchOptions, nil
}
