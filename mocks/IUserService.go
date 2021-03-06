// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/PhillBeck/ingresse-backend/service (interfaces: IUserService)

// Package mock_service is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	golang_odm "github.com/PhillBeck/golang-odm"
	model "github.com/PhillBeck/ingresse-backend/model"
	service "github.com/PhillBeck/ingresse-backend/service"
	gomock "github.com/golang/mock/gomock"
	bson "gopkg.in/mgo.v2/bson"
)

// MockIUserService is a mock of IUserService interface
type MockIUserService struct {
	ctrl     *gomock.Controller
	recorder *MockIUserServiceMockRecorder
}

// MockIUserServiceMockRecorder is the mock recorder for MockIUserService
type MockIUserServiceMockRecorder struct {
	mock *MockIUserService
}

// NewMockIUserService creates a new mock instance
func NewMockIUserService(ctrl *gomock.Controller) *MockIUserService {
	mock := &MockIUserService{ctrl: ctrl}
	mock.recorder = &MockIUserServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIUserService) EXPECT() *MockIUserServiceMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockIUserService) Create(arg0 *model.User) error {
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockIUserServiceMockRecorder) Create(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIUserService)(nil).Create), arg0)
}

// DeleteByID mocks base method
func (m *MockIUserService) DeleteByID(arg0 bson.ObjectId) error {
	ret := m.ctrl.Call(m, "DeleteByID", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByID indicates an expected call of DeleteByID
func (mr *MockIUserServiceMockRecorder) DeleteByID(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByID", reflect.TypeOf((*MockIUserService)(nil).DeleteByID), arg0)
}

// FindByIdAndReplace mocks base method
func (m *MockIUserService) FindByIdAndReplace(arg0 bson.ObjectId, arg1 *model.User) error {
	ret := m.ctrl.Call(m, "FindByIdAndReplace", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// FindByIdAndReplace indicates an expected call of FindByIdAndReplace
func (mr *MockIUserServiceMockRecorder) FindByIdAndReplace(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByIdAndReplace", reflect.TypeOf((*MockIUserService)(nil).FindByIdAndReplace), arg0, arg1)
}

// GetByID mocks base method
func (m *MockIUserService) GetByID(arg0 bson.ObjectId) (*model.User, error) {
	ret := m.ctrl.Call(m, "GetByID", arg0)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID
func (mr *MockIUserServiceMockRecorder) GetByID(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockIUserService)(nil).GetByID), arg0)
}

// Paginate mocks base method
func (m *MockIUserService) Paginate(arg0 service.PaginationOptions) ([]*model.User, *golang_odm.PaginationInfo, error) {
	ret := m.ctrl.Call(m, "Paginate", arg0)
	ret0, _ := ret[0].([]*model.User)
	ret1, _ := ret[1].(*golang_odm.PaginationInfo)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Paginate indicates an expected call of Paginate
func (mr *MockIUserServiceMockRecorder) Paginate(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Paginate", reflect.TypeOf((*MockIUserService)(nil).Paginate), arg0)
}
