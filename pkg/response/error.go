package response

import (
	"errors"
	"net/http"
)

type ErrorRequest struct {
	Code int
	Err  error
}

func (e *ErrorRequest) Error() string {
	return e.Err.Error()
}

func NewErrorRequest(code int, message string) ErrorRequest {
	return ErrorRequest{
		Code: code,
		Err:  errors.New(message),
	}
}

var (
	// User
	InvalidLoginCredentials = NewErrorRequest(http.StatusUnauthorized, "Invalid login credentials")
	UserNotFound            = NewErrorRequest(http.StatusNotFound, "User not found")
	UserAlreadyExists       = NewErrorRequest(http.StatusConflict, "Username or email already exists")

	// JWT
	FailedToGenerateToken = NewErrorRequest(http.StatusInternalServerError, "Failed to generate token")
	FailedToParseToken    = NewErrorRequest(http.StatusUnauthorized, "Failed to parse token")
	InvalidToken          = NewErrorRequest(http.StatusUnauthorized, "Invalid token")

	// Reservation
	ReservationNotFound = NewErrorRequest(http.StatusNotFound, "Reservation not found")
	ReservationExists   = NewErrorRequest(http.StatusConflict, "Reservation already exists")

	// Server
	InternalServerError = NewErrorRequest(http.StatusInternalServerError, "Internal Server Error")
	BadRequest          = NewErrorRequest(http.StatusBadRequest, "Bad Request")
)
