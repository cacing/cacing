package liststruct

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var initialDataSize = uint(0)
var initialData = make([]*Data, initialDataSize)

func TestCreateListStruct(t *testing.T) {
	storage := NewListStruct(initialData)

	if assert.NotNil(t, storage) {
		assert.Equal(t, storage.GetSize(), initialDataSize)
	}
}
