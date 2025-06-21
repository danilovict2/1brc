package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("missing measurements file")
	}

	fmt.Println("Hello, world!")
}
