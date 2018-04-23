package handler_test

import (
	"fmt"
	"testing"

	"github.com/PhillBeck/golang-odm"

	"github.com/PhillBeck/ingresse-backend/lib"
	"github.com/PhillBeck/ingresse-backend/lib/query"
	"github.com/PhillBeck/ingresse-backend/lib/responseHandler"

	"gopkg.in/mgo.v2/bson"

	"github.com/PhillBeck/ingresse-backend/handler"
	"github.com/PhillBeck/ingresse-backend/mocks"
	"github.com/PhillBeck/ingresse-backend/model"
	"github.com/golang/mock/gomock"
)

func TestFindByID(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	userService := mocks.NewMockIUserService(mockCtrl)
	userHandler := handler.User{
		Service: userService}

	// Invalid ID - should return error
	sID := "234453v"
	statusCode, resp := userHandler.FindByID(sID)
	if statusCode != 400 {
		t.Errorf("Expected statusCode 400, got: %d; resp: %+v", statusCode, resp)
	}

	responseUser := model.User{
		Name:     "Phillip",
		CPF:      "999.999.999-99",
		Username: "PhillBeck"}

	// Valid ID, no errors
	sID = "5adce028c0b58a4a19266f4e"
	ID := bson.ObjectIdHex(sID)
	userService.EXPECT().GetByID(ID).Return(&responseUser, nil).Times(1)
	statusCode, resp = userHandler.FindByID(sID)
	if statusCode != 200 {
		t.Errorf("Expected statusCode 200, got: %d; resp: %+v", statusCode, resp)
	}

	// Valid ID, internal error
	userService.EXPECT().GetByID(ID).Return(nil, fmt.Errorf("mongodb error")).Times(1)
	statusCode, resp = userHandler.FindByID(sID)
	if statusCode != 500 {
		t.Errorf("Expected statusCode 500, got: %d; resp: %+v", statusCode, resp)
	}
}

func TestFindByIdAndRemove(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	userService := mocks.NewMockIUserService(mockCtrl)
	userHandler := handler.User{
		Service: userService}

	// Invalid ID - should return error
	sID := "234453v"
	statusCode, resp := userHandler.FindByIdAndRemove(sID)
	if statusCode != 400 {
		t.Errorf("Expected statusCode 400, got: %d; resp: %+v", statusCode, resp)
	}

	// Valid ID, no errors
	sID = "5adce028c0b58a4a19266f4e"
	ID := bson.ObjectIdHex(sID)
	userService.EXPECT().DeleteByID(ID).Return(nil).Times(1)
	statusCode, resp = userHandler.FindByIdAndRemove(sID)
	if statusCode != 204 {
		t.Errorf("Expected statusCode 200, got: %d; resp: %+v", statusCode, resp)
	}

	// Valid ID, internal error
	userService.EXPECT().DeleteByID(ID).Return(fmt.Errorf("mongodb error")).Times(1)
	statusCode, resp = userHandler.FindByIdAndRemove(sID)
	if statusCode != 500 {
		t.Errorf("Expected statusCode 500, got: %d; resp: %+v", statusCode, resp)
	}
}

func TestFindByIdAndReplace(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	userService := mocks.NewMockIUserService(mockCtrl)
	userHandler := handler.User{
		Service: userService}

	validID := "5adce028c0b58a4a19266f4e"
	ID := bson.ObjectIdHex(validID)

	user := model.User{
		Name:     "Phillip",
		CPF:      "999.999.999-99",
		Username: "PhillBeck"}
	user.ID = ID

	// Invalid ID - should return error
	sID := "234453v"
	statusCode, resp := userHandler.FindByIdAndReplace(user, sID)
	if statusCode != 400 {
		t.Errorf("Expected statusCode 400, got: %d; resp: %+v", statusCode, resp)
	}

	// Valid ID, no errors
	userService.EXPECT().FindByIdAndReplace(ID, &user).Return(nil).Times(1)
	statusCode, resp = userHandler.FindByIdAndReplace(user, validID)
	if statusCode != 204 {
		t.Errorf("Expected statusCode 200, got: %d; resp: %+v", statusCode, resp)
	}

	// Valid ID, internal error
	userService.EXPECT().FindByIdAndReplace(ID, &user).Return(fmt.Errorf("mongodb error")).Times(1)
	statusCode, resp = userHandler.FindByIdAndReplace(user, validID)
	if statusCode != 500 {
		t.Errorf("Expected statusCode 500, got: %d; resp: %+v", statusCode, resp)
	}

	// Different ID on body and path
	otherID := bson.NewObjectId()
	statusCode, resp = userHandler.FindByIdAndReplace(user, otherID.Hex())
	if statusCode != 400 {
		t.Errorf("Expected statusCode 400, got: %d; resp: %+v", statusCode, resp)
	}
}

func TestFind(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	userService := mocks.NewMockIUserService(mockCtrl)
	userHandler := handler.User{
		Service: userService}

	// Invalid Query
	statusCode, resp := userHandler.Find("name::", 2, 5)
	if statusCode != 400 || resp.(*lib.ErrorResponse).Message != "Invalid Query" {
		t.Errorf("Expected statusCode 400, got: %d; resp: %+v", statusCode, resp)
	}

	PaginationOptions, _ := query.ParseQuery("name:Phillip", 1, 10)

	// Valid query, internal Error
	userService.EXPECT().Paginate(PaginationOptions).Return(nil, nil, fmt.Errorf("mongoError")).Times(1)
	statusCode, resp = userHandler.Find("name:Phillip", 1, 10)
	if statusCode != 500 {
		t.Errorf("Expected statusCode 500, got: %d; resp: %+v", statusCode, resp)
	}

	// Everything OK
	docs := []*model.User{&model.User{
		Name:     "Phillip",
		CPF:      "999.999.999-99",
		Username: "PhillBeck"}}

	info := &odm.PaginationInfo{
		CurrentPage:    1,
		RecordsPerPage: 10,
		NumPages:       1,
		NumRecords:     1}

	userService.EXPECT().Paginate(PaginationOptions).Return(docs, info, nil).Times(1)
	statusCode, resp = userHandler.Find("name:Phillip", 1, 10)
	tResp := resp.(responseHandler.PaginationResponse)
	if statusCode != 200 ||
		tResp.PerPage != 10 ||
		tResp.Page != 1 ||
		tResp.Pages != 1 ||
		tResp.Total != 1 ||
		tResp.Docs.([]*model.User)[0] != docs[0] {
		t.Errorf("Expected statusCode 500, got: %d; resp: %+v", statusCode, resp)
	}
}
