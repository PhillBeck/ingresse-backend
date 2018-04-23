package handler

import (
	"fmt"
	"net/http"

	"github.com/PhillBeck/ingresse-backend/lib"
	"github.com/PhillBeck/ingresse-backend/model"
	"github.com/PhillBeck/ingresse-backend/service"
	"gopkg.in/macaron.v1"
)

type User struct {
	Service service.IUserService
}

func NewUserHandler() *User {
	return &User{
		Service: service.NewUserService()}
}

func (h *User) GetOne(ctx *macaron.Context) {
	sID := ctx.Params("ID")

	ctx.JSON(h.FindByID(sID))
}

func (h *User) Post(user model.User, ctx *macaron.Context) {
	// Ask for the creation of the document and handles any errors
	err := h.Service.Create(&user)
	if err != nil {
		res := lib.HandleError(err)
		ctx.JSON(res.StatusCode, res)
		return
	}

	// Add header "location"
	ctx.Resp.Header().Add("location", fmt.Sprintf("/users/%s", user.ID.Hex()))
	ctx.JSON(http.StatusCreated, &user)
}

func (h *User) Put(user model.User, ctx *macaron.Context) {
	// Gets path parameter "ID" as a string
	sID := ctx.Params("ID")

	ctx.JSON(h.FindByIdAndReplace(user, sID))
}

func (h *User) Delete(ctx *macaron.Context) {
	// Gets path parameter "ID" as a string
	sID := ctx.Params("ID")

	ctx.JSON(h.FindByIdAndRemove(sID))
}

func (h *User) FindByIdAndRemove(sID string) (int, interface{}) {
	// Check if ID is valid
	ID, err := lib.ParseObjectID(sID)
	if err != nil {
		response := lib.HandleError(err)
		return response.StatusCode, response
	}

	err = h.Service.DeleteByID(ID)
	if err != nil {
		res := lib.HandleError(err)
		return res.StatusCode, res
	}

	return http.StatusNoContent, nil
}

func (h *User) FindByID(sID string) (int, interface{}) {
	ID, err := lib.ParseObjectID(sID)
	if err != nil {
		res := lib.HandleError(err)
		return res.StatusCode, res
	}

	user, err := h.Service.GetByID(ID)
	if err != nil {
		res := lib.HandleError(err)
		return res.StatusCode, res
	}

	return http.StatusOK, user
}

func (h *User) FindByIdAndReplace(user model.User, sID string) (int, interface{}) {
	// Check if ID is valid
	ID, err := lib.ParseObjectID(sID)
	if err != nil {
		response := lib.HandleError(err)
		return response.StatusCode, response
	}

	// This validation is not on the service layer because I judged
	// it has more to do with the formatting of the actual request than
	// with the business logic. It validates if the ID in the path parameter
	// is the same as the ID on the object passed
	if user.GetID() != ID {
		res := lib.HandleError(fmt.Errorf("changing id"))
		return res.StatusCode, res
	}

	// Asks for replacement and handles errors
	err = h.Service.FindByIdAndReplace(ID, &user)
	if err != nil {
		response := lib.HandleError(err)
		return response.StatusCode, response
	}

	// Everything went ok
	return http.StatusNoContent, nil
}
