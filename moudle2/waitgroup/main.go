package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	waitBySleep()
}

func waitBySleep() {
	for i := 0; i < 100; i++ {
		go fmt.Println(i)
	}
	time.Sleep(time.Second)
}

func waiteByChannel() {
	c := make(chan bool, 100)
	for i := 0; i < 100; i++ {
		go func(i int) {
			fmt.Println(i)
			c <- true
		}(i)
	}
	for i := 0; i < 100; i++ {
		<-c
	}
}

func waitByWG() {
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
	wg.Wait()
}
