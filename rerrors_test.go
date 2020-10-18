package rerrors

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBadRequestError(t *testing.T) {

	err := NewBadRequestError(errors.New("Test Error"))

	assert.EqualValues(t, err.Error().Error(), "Test Error")
	assert.EqualValues(t, err.Message(), "Test Error")
	assert.EqualValues(t, err.Status(), http.StatusBadRequest)
	assert.EqualValues(t, len(err.Causes()), 0)
}

func TestNotFoundError(t *testing.T) {

	err := NewNotFoundError(errors.New("Test Error"))

	assert.EqualValues(t, err.Error().Error(), "Test Error")
	assert.EqualValues(t, err.Message(), "Test Error")
	assert.EqualValues(t, err.Status(), http.StatusNotFound)
	assert.EqualValues(t, len(err.Causes()), 0)
}

func TestUnauthorizedError(t *testing.T) {

	err := NewUnauthorizedError(errors.New("Test Error"))

	assert.EqualValues(t, err.Error().Error(), "Test Error")
	assert.EqualValues(t, err.Message(), "Test Error")
	assert.EqualValues(t, err.Status(), http.StatusUnauthorized)
	assert.EqualValues(t, len(err.Causes()), 0)
}

func TestInternalServerError(t *testing.T) {

	err := NewInternalServerError(errors.New("Test Error"))

	assert.EqualValues(t, err.Error().Error(), "Test Error")
	assert.EqualValues(t, err.Message(), "Test Error")
	assert.EqualValues(t, err.Status(), http.StatusInternalServerError)
	assert.EqualValues(t, len(err.Causes()), 0)
}
