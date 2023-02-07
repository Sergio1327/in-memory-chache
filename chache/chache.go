package chache

import (
	"errors"
	"sync"
	"time"
)

var mutex sync.RWMutex

type Chache struct {
	chache map[string]interface{}
}

func New() *Chache {
	return &Chache{chache: make(map[string]interface{})}
}

func (c Chache) Set(key string, value interface{}, t time.Duration) error {
	mutex.Lock()
	c.chache[key] = value
	mutex.Unlock()
	go func() {
		time.Sleep(t)
		c.Delete(key)
	}()
	return nil
}

func (c Chache) Get(key string) interface{} {
	mutex.RLock()
	defer mutex.RUnlock()
	_, ok := c.chache[key]
	if !ok {
		return errors.New("kэш с таким ключом не найден")
	}
	return c.chache[key]
}

func (c Chache) Delete(key string) error {
	mutex.Lock()
	defer mutex.Unlock()
	_, ok := c.chache[key]
	if !ok {
		return errors.New("kэш с таким ключом не найден")
	}
	delete(c.chache, key)
	return nil
}
