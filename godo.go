package main

import (
	"fmt"

	"github.com/tracyde/godo/collection"
)

func main() {
	c := collection.New("/tmp/test")

	fmt.Printf("Type: %T  ::  Value: %v\n", c, c)
}
