package chache

import (
	"errors"
	"sync"
	"time"
)

type Chache struct {
	mu     sync.Mutex
	chache map[string]interface{}
}

func New() *Chache {
	return &Chache{chache: make(map[string]interface{})}
}

func (c Chache) Set(key string, value interface{}, t time.Duration) error {
	c.mu.Lock()
	c.chache[key] = value
	c.mu.Unlock()
	go func() {
		time.Sleep(t)
		c.Delete(key)
	}()
	return nil
}

func (c Chache) Get(key string) interface{} {
	c.mu.Lock()
	defer c.mu.Unlock()
	_, ok := c.chache[key]
	if !ok {
		return errors.New("kэш с таким ключом не найден")
	}
	return c.chache[key]
}

func (c Chache) Delete(key string) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	_, ok := c.chache[key]
	if !ok {
		return errors.New("kэш с таким ключом не найден")
	}
	delete(c.chache, key)
	return nil
}
