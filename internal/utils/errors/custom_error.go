package errors

import (
	"fmt"
	"net/http"
)

var (
	// general
	ErrInvalidBodyPayload = NewCustomError(BadRequest, "Invalid Body Payload")
	ErrInternalServer     = NewCustomError(InternalServer, http.StatusText(http.StatusInternalServerError))
	ErrInvalidAuthHeader  = NewCustomError(BadRequest, "error invalid auth header")

	// JWT
	ErrAccessTokenExpired = NewCustomError(Unauthorized, "Token has already expired")
	ErrInvalidToken       = NewCustomError(Unauthorized, "Token is invalid")

	// User
	ErrRecordNotFound  = NewCustomError(BadRequest, "error invalid auth header")
)

func GenerateErrQueryParamRequired(param string) *CustomError {

	return NewCustomError(BadRequest, fmt.Sprintf("Query param %s is required", param))
}

func GenerateErrQueryParamInvalid(param string) *CustomError {

	return NewCustomError(BadRequest, fmt.Sprintf("Query param %s is invalid", param))
}

func GenerateErrPathParamInvalid(param string) *CustomError {

	return NewCustomError(BadRequest, fmt.Sprintf("Path param %s is invalid", param))
}
