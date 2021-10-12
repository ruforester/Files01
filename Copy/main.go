package main

import (
	"io"
	"log"
	"os"
)

func main() {
	originalFile, err := os.Open("text.txt")
	if err != nil {
		log.Fatal(err)
	}

	newFile, err := os.Create("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	bytesWritten, err := io.Copy(newFile, originalFile)
	if err != nil {
		log.Fatal(err)
	}

	err = newFile.Truncate(100)
	if err != nil {
		log.Fatal(err)
	}
	newFile.Close()
	originalFile.Close()

	log.Printf("Copied %d bytes", bytesWritten)

	err = os.Remove(originalFile.Name())
	if err != nil {
		log.Fatal(err)
	}

}
