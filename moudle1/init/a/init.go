package a

import (
	"fmt"

	_ "github.com/liushiju/golang_examples/module1/init/b"
)

func init() {
	fmt.Println("init from a")
}
