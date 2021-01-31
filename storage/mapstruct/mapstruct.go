package mapstruct

import (
	"fmt"
	"sync"
	"time"

	"github.com/cacing/cacing/storage"
)

// Data struct define structure of saved data
type Data struct {
	Value      interface{}
	Expiration int64
}

// MapStruct is storage engine that store data into list of structs
type MapStruct struct {
	Data map[string]Data
	Mu   sync.RWMutex
}

// NewMapStruct generate new MapStruct
func NewMapStruct(initialData map[string]Data) storage.Storage {
	return &MapStruct{
		Data: initialData,
	}
}

// GetSize return how many datum in storage
func (store *MapStruct) GetSize() uint {
	store.Mu.RLock()
	defer store.Mu.RUnlock()
	return uint(len(store.Data))
}

// Exists return true if data exists
// or false if doesn't
func (store *MapStruct) Exists(key string) bool {
	store.Mu.RLock()
	defer store.Mu.RUnlock()

	// check data expiration
	data, exist := store.Data[key]
	if exist {
		if data.Expiration > 0 {
			if time.Now().UnixNano() > data.Expiration {
				delete(store.Data, key)
				return false
			}
		}
		return true
	}

	return false
}

// Get return value with inserted key
// if this key doesn't exists, return error
func (store *MapStruct) Get(key string) (value interface{}, err error) {
	store.Mu.RLock()
	defer store.Mu.RUnlock()

	// check data existence
	exist := store.Exists(key)
	if !exist {
		return nil, fmt.Errorf("data not found")
	}

	// get data after cleaned
	data := store.Data[key]

	return data.Value, nil
}

// Set to add data into storage
// and return error if any problems happen
func (store *MapStruct) Set(key string, val interface{}, t time.Duration) error {

	store.Mu.Lock()
	var ex int64
	if t > 0 {
		ex = time.Now().Add(t).UnixNano()
	}

	store.Data[key] = Data{
		Value:      val,
		Expiration: ex,
	}
	store.Mu.Unlock()

	return nil
}

func (store *MapStruct) SetExpired(key string, t time.Duration) (interface{}, error) {

	store.Mu.Lock()

	data, exist := store.Data[key]
	if !exist {
		return nil, fmt.Errorf("data not found")
	}

	data.Expiration = time.Now().Add(t).UnixNano()
	store.Data[key] = data

	store.Mu.Unlock()

	return data, nil
}

// Delete return deleted value
// or error if any problems happen
func (store *MapStruct) Delete(key string) (interface{}, error) {
	// check data existence
	store.Mu.Lock()

	data, err := store.Get(key)
	if err != nil {
		return nil, fmt.Errorf("data not found")
	}

	delete(store.Data, key)

	store.Mu.Unlock()
	return data, nil
}
