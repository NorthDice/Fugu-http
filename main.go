package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func getLinesChannel(f io.ReadCloser) <-chan string {
	out := make(chan string, 1)

	go func() {
		defer close(out)
		defer f.Close()
		buf := make([]byte, 16)

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
					out <- lineAccum.String()
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
			out <- lineAccum.String()
		}
	}()

	return out
}

func main() {

	f, err := os.Open("messages.txt")
	if err != nil {
		log.Fatalf("Cannot open the file, err = %v", err)
	}
	defer f.Close()

	lines := getLinesChannel(f)
	for line := range lines {
		fmt.Printf("read: %s\n", line)
	}
}
