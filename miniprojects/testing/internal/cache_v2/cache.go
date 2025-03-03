package cache

import (
	"errors"
)

var ErrNotFound = errors.New("value not found")

type Cache interface {
	Set(key, value string) error
	Get(key string) (string, error)
	Delete(key string) error
}

type CacheImpl struct{
	contaier map[string]string
}

func NewCache() Cache {
	return &CacheImpl{
		contaier: make(map[string]string),
	}
}

func (c *CacheImpl) Set(key, value string) error {
	c.contaier[key] = value
	return nil
}

func (c *CacheImpl) Get(key string) (string, error) {
	if val, ok := c.contaier[key]; ok {
		return val, nil
	} else {
		return string(""), ErrNotFound
	}
}

func (c *CacheImpl) Delete(key string) error {
	delete(c.contaier, key)
	return nil
}