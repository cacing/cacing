package liststruct

import (
	"fmt"
	"time"

	"github.com/needkopi/cacing/storages"
)

// Data struct define structure of saved data
type Data struct {
	Value       interface{}
	Expiration int64
}

// ListStruct is storage engine that store data into list of structs
type ListStruct struct {
	Data map[string]Data
}

// NewListStruct generate new ListStruct
func NewListStruct(initialData map[string]Data) storages.Storage {
	return &ListStruct{
		Data: initialData,
	}
}

// GetSize return how many datum in storage
func (store *ListStruct) GetSize() uint {
	return uint(len(store.Data))
}

// Get return value with inserted key
// if this key doesn't exists, return error
func (store *ListStruct) Get(key string) (value interface{},err error) {

	val, exist := store.Data[key]
	if !exist {
		return nil, fmt.Errorf("data not found")
	}

	if val.Expiration > 0 {
		if time.Now().UnixNano() > val.Expiration {
			return nil, fmt.Errorf("data not found")
		}
	}

	return val.Value, nil

}

// Set to add data into storage
// and return error if any problems happen
func (store *ListStruct) Set(key string, val interface{}, t time.Duration) error {

	var ex int64
	if t > 0 {
		ex = time.Now().Add(t).UnixNano()
	}

	store.Data[key] = Data{
		Value:        val,
		Expiration: ex,
	}

	return nil
}
