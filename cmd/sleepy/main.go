package main

import (
	"fmt"

	"github.com/joshkim456/mc-control-plane/internal/properties"
)

func main() {
	fmt.Println("sleepy: starting")
	fmt.Println(properties.Parse())
}