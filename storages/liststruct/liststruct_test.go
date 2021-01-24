package liststruct

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var initialDataSize = uint(0)
var initialData = map[string]Data{

var key = "id"
var val = uint(1238782738973)

func TestCreateListStruct(t *testing.T) {
	storage := NewListStruct(initialData)

	if assert.NotNil(t, storage) {
		assert.Equal(t, storage.GetSize(), initialDataSize)
	}
}

func TestSetToListStruct(t *testing.T) {
	storage := NewListStruct(initialData)
	storage.Set(key, val, 12 * time.Second)

	assert.Equal(t, storage.GetSize(), uint(1))
}

func TestGetFromListStruct(t *testing.T) {
	storage := NewListStruct(initialData)
	storage.Set(key, val, 0)

	data, err := storage.Get(key)
	if assert.Nil(t, err) {
		assert.Equal(t, data, val)
	}
}
