package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("text.txt", os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(file)
	file.Close()
}
