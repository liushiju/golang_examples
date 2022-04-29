package main

import (
	"fmt"
	_ "https://github.com/liushiju/golang_examples/init/a"
	_ "https://github.com/liushiju/golang_examples/init/b"
)

func init()  {
	fmt.Println("main init")
}

func main()  {
	fmt.Println("main function")
}