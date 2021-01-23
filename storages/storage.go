package storages

// Storage abstraction
// all storage types should follow this interface
type Storage interface {
	Set(key string, val interface{}) error
	Get(key string) (interface{}, error)

	GetSize() uint
}
