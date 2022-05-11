package main

import "fmt"

func main()  {
	ch := make(chan int)
	go func() {
		fmt.Println("from child thread")
		ch <- 1

	}()
	<- ch
}