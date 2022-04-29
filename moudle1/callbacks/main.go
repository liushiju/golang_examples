package main

import . "fmt"

func main() {
	DoOperation(1, increase)
	DoOperation(1, decrease)
}

func increase(a, b int) {
	Println("increase result is:", a+b)
}

func decrease(a, b int) {
	Println("decrease resule is:", a-b)
}

func DoOperation(y int, f func(int, int)) {
	f(y, 1)
}
