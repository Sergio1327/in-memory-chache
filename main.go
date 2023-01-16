package main

import (
	"fmt"
	"modules/chache"
)

func main() {
	a := chache.New()
	a.Set("123", "An Sergey Aleksandrovich")
	a.Get("123")
	a.Delete("123")
	b:=chache.New()
	fmt.Println(b)
}
