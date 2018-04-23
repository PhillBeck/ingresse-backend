package query_test

import (
	"fmt"
	"testing"

	"github.com/PhillBeck/ingresse-backend/lib/query"
	"github.com/PhillBeck/ingresse-backend/service"
	"gopkg.in/mgo.v2/bson"
)

func TestParseQuery(t *testing.T) {
	// All valid
	if err := allValid(); err != nil {
		t.Errorf("%s", err.Error())
	}

	if err := invalidQuery(); err != nil {
		t.Errorf("%s", err.Error())
	}

	if err := invalidPage(); err != nil {
		t.Errorf("%s", err.Error())
	}
}

func allValid() error {
	opt, err := query.ParseQuery("name:Phillip", 2, 5)
	if err != nil {
		return err
	}

	expected := service.PaginationOptions{
		Page:           2,
		RecordsPerPage: 5,
		Query:          bson.M{"name": bson.M{"$regex": "Phillip"}}}

	if opt.Page != expected.Page ||
		opt.RecordsPerPage != expected.RecordsPerPage ||
		opt.Query["name"].(bson.M)["$regex"] != expected.Query["name"].(bson.M)["$regex"] {
		return fmt.Errorf("AllValid Error! Expected: %+v, got: %+v", expected, opt)
	}

	return nil
}

func invalidQuery() error {
	expected := "Invalid Query"

	_, err := query.ParseQuery("name::", 2, 5)
	if err == nil {
		return fmt.Errorf("InvalidQuery Error! Expected: %s, got: nil", expected)
	}

	if err.Error() != expected {
		return fmt.Errorf("InvalidQuery Error! Expected: %s, got: %s", expected, err.Error())
	}

	return nil
}

func invalidPage() error {
	expected := service.PaginationOptions{
		Page:           1,
		RecordsPerPage: 10}

	opt, err := query.ParseQuery("", 0, -2)
	if err != nil {
		return fmt.Errorf("InvalidPage Error! Expected: nil, got: %s", err.Error())
	}

	if opt.Page != expected.Page || opt.RecordsPerPage != expected.RecordsPerPage {
		return fmt.Errorf("InvalidPage Error! Expected: %+v, got: %+v", expected, opt)
	}

	return nil
}
