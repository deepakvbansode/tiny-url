package errors

import (
	"golang.frontdoorhome.com/personal-project/tiny-url/internal/literals"
	"net/http"
)

const (
	envPrefix = "SRFACADE"
)

type ServerError struct {
	Message          string
	Code             string
	HTTPResponseCode int
}

func (err *ServerError) Error() string {
	return err.Code + ": " + err.Message
}

var EmptyObjectIdentifierError = ServerError{
	Message:          literals.EmptyIdentifier,
	Code:             envPrefix + "_F_0001",
	HTTPResponseCode: http.StatusInternalServerError,
}

var ObjectAlreadyExistsInFactoryError = ServerError{
	Message:          literals.ObjectAlreadyExist,
	Code:             envPrefix + "_F_0001",
	HTTPResponseCode: http.StatusInternalServerError,
}

var InvalidObjectIdentifierError = ServerError{
	Message:          literals.InvalidObjectIdentifier,
	Code:             envPrefix + "_F_0001",
	HTTPResponseCode: http.StatusInternalServerError,
}
