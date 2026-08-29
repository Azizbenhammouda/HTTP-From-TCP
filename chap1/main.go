package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func getLinesChannel(f io.ReadCloser) <-chan string {
	ch := make(chan string)
	go func() {
		defer f.Close()
		defer close(ch)

		LINE := ""

		for {
			buf := make([]byte, 8)

			n, err := f.Read(buf)
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal(err)
			}

			parts := strings.Split(string(buf[:n]), "\n")

			for i := 0; i < len(parts)-1; i++ {
				ch <- (LINE + parts[i])
				LINE = ""
			}

			LINE += parts[len(parts)-1]
		}

		if LINE != "" {
			ch <- LINE
		}
	}()

	return ch
}

func main() {
	f, err := os.Open("messages.txt")
	if err != nil {
		log.Fatal(err)
	}

	for line := range getLinesChannel(f) {
		fmt.Printf("read: %s\n", line)
	}
}
