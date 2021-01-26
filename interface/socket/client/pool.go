package client

import (
	"errors"
	"sync"

	uuid "github.com/satori/go.uuid"
)

// Pool type to store connected clients
type Pool struct {
	IDs []uuid.UUID
	mu  sync.RWMutex
}

// NewPool create a new client pool
func NewPool() *Pool {
	return &Pool{}
}

// Add func append a new client into pool
func (p *Pool) Add() (uuid.UUID, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	newID := uuid.NewV4()
	p.IDs = append(p.IDs, newID)
	return newID, nil
}

// IsExists check is the given id registered in client pool
func (p *Pool) IsExists(id string) (bool, int) {
	p.mu.RLock()
	defer p.mu.RUnlock()

	index := -1
	idParsed := uuid.FromStringOrNil(id)
	if uuid.Equal(idParsed, uuid.Nil) {
		return false, index
	}

	for i, idInPool := range p.IDs {
		if uuid.Equal(idInPool, idParsed) {
			index = i
			return true, index
		}
	}

	return false, index
}

// Delete func remove client from client pool
func (p *Pool) Delete(id string) (uuid.UUID, error) {
	p.mu.RLock()
	defer p.mu.RUnlock()

	exists, index := p.IsExists(id)
	if !exists {
		return uuid.Nil, errors.New("invalid id")
	}

	p.IDs = append(p.IDs[:index], p.IDs[index+1:]...)

	return uuid.FromStringOrNil(id), nil
}
