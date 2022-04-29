package a

import (
	"fmt"

_ "https://github.com/liushiju/golang_examples/init/b"
)

func init() {
	fmt.Println("init from a")
}
