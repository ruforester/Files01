package main

import (
	"log"
	"os"
)

func main() {
	originalPath := "text.txt"
	newPath := "text1.txt"
	err := os.Rename(originalPath, newPath)
	if err != nil {
		log.Fatal(err)
	}
}
