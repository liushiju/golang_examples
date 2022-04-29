package main

import (
	"fmt"

	_ "github.com/liushiju/golang_examples/moudle1/init/a"
	_ "github.com/liushiju/golang_examples/moudle1/init/b"
)

func init() {
	fmt.Println("main init")
}

func main() {
	fmt.Println("main function")
}
