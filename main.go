package main

import (
	"fmt"
	"log"
	"os"
)

// compile passing -ldflags "-X main.Build <build sha1>"
var Build string

func main() {
	fmt.Printf("Using build: %s\n", Build)
}
