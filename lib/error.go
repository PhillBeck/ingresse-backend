package lib

import (
	"net/http"
)

type ErrorResponse struct {
	StatusCode int    `json:"statusCode"`
	Type       string `json:"type"`
	Message    string `json:"message"`
}

func HandleError(err error) *ErrorResponse {
	switch err.Error() {
	case "not found":
		return handleNotFound(err)
	case "invalid id":
		return handleInvalidID(err)
	case "changing id":
		return handleChangingID()
	case "Invalid Query":
		return handleInvalidQuery()
	default:
		return handleInternal(err)
	}
}

func handleInvalidQuery() *ErrorResponse {
	return &ErrorResponse{
		StatusCode: 400,
		Type:       "Validation Error",
		Message:    "Invalid Query"}
}

func handleNotFound(err error) *ErrorResponse {
	return &ErrorResponse{
		StatusCode: 404,
		Type:       "Not Found",
		Message:    "Not Found"}
}

func handleInternal(err error) *ErrorResponse {
	return &ErrorResponse{
		StatusCode: 500,
		Type:       "Internal Server Error",
		Message:    "Internal server error, please try again in a few minutes."}
}

func handleInvalidID(err error) *ErrorResponse {
	return &ErrorResponse{
		StatusCode: http.StatusBadRequest,
		Type:       "Invalid ID",
		Message:    "IDs should be a 24 characters long hexadecimal value"}
}

func handleChangingID() *ErrorResponse {
	return &ErrorResponse{
		StatusCode: http.StatusBadRequest,
		Type:       "Validation Error",
		Message:    "IDs from body and path do not match. Changing ID is not allowed"}
}
