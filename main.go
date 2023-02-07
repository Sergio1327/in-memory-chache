package main

import (
	"fmt"
	"github.com/Sergio1327/in-memory-chache/chache"
	"time"
)

func main() {
	a := chache.New()
	a.Set("qwerty", "AnSergeyAleksandrovich", time.Second*5)
	time.Sleep(time.Second * 6)
	fmt.Println(a.Get("qwerty"))
	a.Delete("qwerty")

}
