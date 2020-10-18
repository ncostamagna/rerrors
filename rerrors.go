package rerrors

import (
	"net/http"
)

// RestErr Object
type RestErr interface {
	Message() string
	Status() int
	Error() error
	Causes() []interface{}
}

type restErr struct {
	ErrMessage string
	ErrStatus  int
	ErrError   error
	ErrCauses  []interface{}
}

func (e restErr) Error() error {
	return e.ErrError
}

func (e restErr) Message() string {
	return e.ErrMessage
}

func (e restErr) Status() int {
	return e.ErrStatus
}

func (e restErr) Causes() []interface{} {
	return e.ErrCauses
}

// NewBadRequestError - create bad request error
func NewBadRequestError(err error) RestErr {
	return restErr{
		ErrMessage: err.Error(),
		ErrStatus:  http.StatusBadRequest,
		ErrError:   err,
	}
}

// NewNotFoundError - create not found error
func NewNotFoundError(err error) RestErr {
	return restErr{
		ErrMessage: err.Error(),
		ErrStatus:  http.StatusNotFound,
		ErrError:   err,
	}
}

// NewUnauthorizedError - create unauthorized error
func NewUnauthorizedError(err error) RestErr {
	return restErr{
		ErrMessage: err.Error(),
		ErrStatus:  http.StatusUnauthorized,
		ErrError:   err,
	}
}

// NewInternalServerError - create internal server error
func NewInternalServerError(err error) RestErr {
	result := restErr{
		ErrMessage: err.Error(),
		ErrStatus:  http.StatusInternalServerError,
		ErrError:   err,
	}

	return result
}
