package liststruct

import (
	"cacing/storages"
	"fmt"
)

// Data struct define structure of saved data
type Data struct {
	Key string
	Val interface{}
}

// ListStructStorage is storage engine that store data into list of structs
type ListStructStorage struct {
	Data []*Data
}

// NewListStructStorage generate new ListStructStorage
func NewListStructStorage(initialData []*Data) storages.Storage {
	return &ListStructStorage{
		Data: initialData,
	}
}

// GetSize return how many datum in storage
func (store ListStructStorage) GetSize() uint {
	return uint(len(store.Data))
}

// Get return value with inserted key
// if this key doesn't exists, return error
func (store ListStructStorage) Get(key string) (interface{}, error) {
	var val interface{}
	found := false
	for _, data := range store.Data {
		if data.Key == key {
			val = data.Val
			found = true
			break
		}
	}

	if !found {
		return nil, fmt.Errorf("No data with key `%s` found", key)
	}

	return val, nil
}

// Set to add data into storage
// and return error if any problems happen
func (store *ListStructStorage) Set(key string, val interface{}) error {
	newData := &Data{
		Key: key,
		Val: val,
	}
	store.Data = append(store.Data, newData)
	return nil
}
