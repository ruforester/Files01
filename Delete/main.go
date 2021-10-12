package main

import (
	"log"
	"os"
)

func main() {
	err := os.Remove("text1.txt")
	if err != nil {
		log.Fatal(err)
	}
}
