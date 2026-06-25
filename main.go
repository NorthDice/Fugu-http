package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	f, err := os.Open("messages.txt")
	if err != nil {
		log.Fatalf("Cannot open the file, err = %v", err)
	}
	defer f.Close()

	buf := make([]byte, 32)

	var lineAccum bytes.Buffer
	for {
		n, err := f.Read(buf)
		if n > 0 {
			currentData := buf[:n]

			for {
				idx := bytes.IndexByte(currentData, '\n')
				if idx == -1 {
					lineAccum.Write(currentData)
					break
				}

				lineAccum.Write(currentData[:idx])
				fmt.Printf("read: %s\n", lineAccum.String())

				lineAccum.Reset()

				currentData = currentData[idx+1:]
			}
		}
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("Read error: %v", err)
		}

	}
	if lineAccum.Len() > 0 {
		fmt.Printf("read: %s\n", lineAccum.String())
	}

}
