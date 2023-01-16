package chache

import (
	"errors"
)

type Chache struct {
	chache map[string]interface{}
}

func New() *Chache {
	return &Chache{chache: make(map[string]interface{})}
}

func (c Chache) Set(key string, value interface{}) error {
	c.chache[key] = value
	_, ok := c.chache[key]
	if !ok {
		return errors.New("kэш с таким ключом не найден")
	}
	return nil
}

func (c Chache) Get(key string) interface{} {
	_, ok := c.chache[key]
	if !ok {
		return errors.New("kэш с таким ключом не найден")
	}
	return c.chache[key]
}

func (c Chache) Delete(key string) error {
	_, ok := c.chache[key]
	if !ok {
		return errors.New("kэш с таким ключом не найден")
	}
	delete(c.chache, key)
	return nil
}
