package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("messages.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

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
			fmt.Printf("read: %s%s\n", LINE, parts[i])
			LINE = ""
		}

		LINE += parts[len(parts)-1]
	}

	if LINE != "" {
		fmt.Printf("read: %s\n", LINE)
	}
}
