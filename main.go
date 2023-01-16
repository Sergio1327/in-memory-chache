package main

import (
	"github.com/Sergio1327/in-memory-chache/chache"
)

func main() {
	a:=chache.New()
	a.Set("123", "An Sergey Aleksandrovich")
	a.Get("123")
	a.Delete("123")
}
