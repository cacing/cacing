package mapstruct

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var initialDataSize = uint(0)
var initialData = map[string]Data{}

var key = "id"
var val = uint(1238782738973)

func TestCreateMapStruct(t *testing.T) {
	storage := NewMapStruct(initialData)

	if assert.NotNil(t, storage) {
		assert.Equal(t, storage.GetSize(), initialDataSize)
	}
}

func TestSetToMapStruct(t *testing.T) {
	storage := NewMapStruct(initialData)
	storage.Set(key, val, 12*time.Second)

	assert.Equal(t, storage.GetSize(), uint(1))
}

func TestGetFromMapStruct(t *testing.T) {
	storage := NewMapStruct(initialData)
	storage.Set(key, val, 0)

	data, err := storage.Get(key)
	if assert.Nil(t, err) {
		assert.Equal(t, data, val)
	}
}

func TestDeleteFromMapStruct(t *testing.T) {
	storage := NewMapStruct(initialData)
	storage.Set(key, val, 0)

	deleted, err := storage.Delete(key)
	if assert.Nil(t, err) {
		assert.Equal(t, deleted, val)

		_, err := storage.Get(key)
		assert.NotNil(t, err)
	}
}
