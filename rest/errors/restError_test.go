package errors

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewInternalServerError(t *testing.T) {
	err := NewInternalServerError("my error message", errors.New("my error"))
	assert.NotNil(t, err)
	assert.EqualValues(t, err.Code, 500)
	assert.EqualValues(t, err.Message, "my error message")
	assert.EqualValues(t, err.Message, err.Error())
	assert.NotNil(t, err.Causes)
	assert.EqualValues(t, 1, len(err.Causes))
	assert.EqualValues(t, "my error", err.Causes[0].(error).Error())
}

func TestNewBadRequestError(t *testing.T) {
	err := NewBadRequestError("my error message")
	assert.NotNil(t, err)
	assert.EqualValues(t, err.Code, 400)
	assert.EqualValues(t, err.Message, "my error message")
	assert.EqualValues(t, err.Message, err.Error())
}

func TestNewNotFoundError(t *testing.T) {
	err := NewNotFoundError("my error message")
	assert.NotNil(t, err)
	assert.EqualValues(t, err.Code, 404)
	assert.EqualValues(t, err.Message, "my error message")
	assert.EqualValues(t, err.Message, err.Error())
}

func TestNewDatabaseError(t *testing.T) {
	err := NewDatabaseError()
	assert.NotNil(t, err)
	assert.EqualValues(t, err.Code, 500)
	assert.EqualValues(t, err.Message, "database error")
	assert.EqualValues(t, err.Message, err.Error())
}

func TestNewUnauthorizedError(t *testing.T) {
	err := NewUnauthorizedError("request requires authorization")
	assert.NotNil(t, err)
	assert.EqualValues(t, err.Code, 401)
	assert.EqualValues(t, err.Message, "request requires authorization")
	assert.EqualValues(t, err.Message, err.Error())
}
