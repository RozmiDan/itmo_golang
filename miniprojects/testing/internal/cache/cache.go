package cache

import (
	"errors"
	"sync"
)

var ErrNotFound = errors.New("value not found")

type Cache interface {
	Set(key, value string) error
	Get(key string) (string, error)
	Delete(key string) error
}

type CacheImpl struct{
	contaier map[string]string
	mx *sync.RWMutex
}

func NewCache() Cache {
	return &CacheImpl{
		contaier: make(map[string]string),
		mx: &sync.RWMutex{},
	}
}

func (c *CacheImpl) Set(key, value string) error {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.contaier[key] = value
	return nil
}

func (c *CacheImpl) Get(key string) (string, error) {
	c.mx.RLock()
	defer c.mx.RUnlock()
	if val, ok := c.contaier[key]; ok {
		return val, nil
	} else {
		return string(""), ErrNotFound
	}
}

func (c *CacheImpl) Delete(key string) error {
	c.mx.Lock()
	defer c.mx.Unlock()
	delete(c.contaier, key)
	return nil
}