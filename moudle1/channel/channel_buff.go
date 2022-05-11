package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := make(chan int, 10)
	var i int = 0
	go func() {
		for i < 10 {
			rand.Seed(time.Now().UnixNano())
			n := rand.Intn(10)
			fmt.Println("putting:", n)
			ch <- n
			i++
		}
		close(ch)
	}()
	fmt.Println("hello from main")

	for v := range ch {
		fmt.Println("receiving: ", v)
	}
}
