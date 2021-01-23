package storages

import "time"

// Storage abstraction
// all storage types should follow this interface
type Storage interface {
	Set(key string, val interface{}, t time.Duration) error
	Get(key string) (value interface{},err error)
	Delete(key string) (interface{}, error)
	Update(k string, v interface{}) (interface{}, error)

	GetSize() uint
}
