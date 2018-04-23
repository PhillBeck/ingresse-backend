package service_test

import (
	"fmt"
	"testing"

	"gopkg.in/mgo.v2/bson"

	"github.com/PhillBeck/ingresse-backend/mocks"
	"github.com/PhillBeck/ingresse-backend/model"
	"github.com/PhillBeck/ingresse-backend/service"
	"github.com/golang/mock/gomock"
)

func TestSaveUser(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	userRepo := mocks.NewMockIUser(mockCtrl)

	userService := service.User{
		Repository: userRepo}

	user := model.User{
		Name:     "Phillip",
		CPF:      "999.999.999-99",
		Username: "PhillBeck"}

	userRepo.EXPECT().Save(&user).Return(nil).Times(1)

	err := userService.Create(&user)
	if err != nil {
		t.Errorf("%s", err.Error())
	}
}

func TestReplaceUser(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	userRepo := mocks.NewMockIUser(mockCtrl)
	userService := service.User{
		Repository: userRepo}

	// Should return invalid ID
	if err := replace_invalidID(&userService); err != nil {
		t.Errorf("%s", err.Error())
	}

	replace_notFound(&userService, userRepo)
	replace_OK(&userService, userRepo)

}

func replace_invalidID(s *service.User) error {
	// Sending user without ID
	user := model.User{
		Name:     "Phillip",
		CPF:      "999.999.999-99",
		Username: "PhillBeck"}

	err := s.FindByIdAndReplace(bson.NewObjectId(), &user)
	if err == nil || err.Error() != "Cannot change property ID" {
		return fmt.Errorf("Did not get invalid ID message")
	}

	return nil
}

func replace_notFound(s *service.User, mockRepo *mocks.MockIUser) {
	user := &model.User{
		Name:     "Phillip",
		CPF:      "999.999.999-99",
		Username: "PhillBeck"}

	user.ID = bson.NewObjectId()

	mockRepo.EXPECT().GetByID(user.ID).Return(nil, fmt.Errorf("not found")).Times(1)

	s.FindByIdAndReplace(user.ID, user)
}

func replace_OK(s *service.User, mockRepo *mocks.MockIUser) {
	user := &model.User{
		Name:     "Phillip",
		CPF:      "999.999.999-99",
		Username: "PhillBeck"}

	user.ID = bson.NewObjectId()

	mockRepo.EXPECT().GetByID(user.ID).Return(user, nil).Times(1)
	mockRepo.EXPECT().Save(user).Return(nil).Times(1)

	s.FindByIdAndReplace(user.ID, user)
}
