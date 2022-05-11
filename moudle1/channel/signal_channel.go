package main

func main() {
	var c = make(chan int)
	go prod(c)
	go comsume(c)
}

func prod(ch chan<- int) {
	for {
		ch <- 1
	}
}

func comsume(ch <-chan int) {
	for {
		<-ch
	}
}
