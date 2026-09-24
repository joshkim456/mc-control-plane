package main

import (
	"fmt"
	"github.com/joshkim456/mc-control-plane/internal/properties"
	"os"
)

func main() {
	fmt.Println("sleepy: starting")

	file, fileOpenErr := os.Open("testdata/server.properties")
	if fileOpenErr != nil {
		fmt.Println("Error:", fileOpenErr)
		os.Exit(1)
	}
	defer file.Close()

	props, parseErr := properties.Parse(file)

	if parseErr != nil {
		fmt.Println("Error:", parseErr)
		os.Exit(1)
	}

	maxPlayers, ok := props.Get("max-players")
	fmt.Println(maxPlayers, ok)
}
