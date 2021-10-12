package main

import (
	"log"
	"os"
)

func main() {
	fileInfo, err := os.Stat("text.txt")
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("File does not exist!")
		}
		log.Fatal(err)
	}

	log.Println("File exists. File info:")
	log.Println(fileInfo)

}
