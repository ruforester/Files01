package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("test.txt", os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	bufferedWriter := bufio.NewWriter(file)
	bytesWritten, err := bufferedWriter.Write([]byte{65, 66, 67})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Bytes written: %d\n", bytesWritten)

	bytesWritten, err = bufferedWriter.WriteString("Buffered string\n")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Bytes written (string): %d\n", bytesWritten)

	unflushedBufferSize := bufferedWriter.Buffered()
	log.Printf("Buffered: %d\n", unflushedBufferSize)
	bytesAvaliable := bufferedWriter.Available()
	log.Printf("Available: %d\n", bytesAvaliable)

	bufferedWriter.Flush()
	bytesAvaliable = bufferedWriter.Available()
	log.Printf("Available: %d\n", bytesAvaliable)

	bufferedWriter.Reset(bufferedWriter)
	bytesAvaliable = bufferedWriter.Available()
	log.Printf("Available: %d\n", bytesAvaliable)

	bufferedWriter = bufio.NewWriterSize(bufferedWriter, 8000)
	bytesAvaliable = bufferedWriter.Available()
	log.Printf("Available: %d\n", bytesAvaliable)

}
