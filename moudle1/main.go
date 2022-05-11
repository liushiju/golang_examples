package main

import (
	"fmt"
	"time"
)

func produce(ch chan int, slice []int) {
	// defer close(ch)
	for _, v := range slice {
		fmt.Println("produce:", v)
		ch <- v
		time.Sleep(1 * time.Second)
	}

}

func consume(ch chan int) {
	for v := range ch {
		fmt.Println("consume:", v)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	ch := make(chan int, 10)
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	go produce(ch, slice)
	// time.Sleep(10 * time.Second)
	go consume(ch)
	time.Sleep(10 * time.Second)
}
