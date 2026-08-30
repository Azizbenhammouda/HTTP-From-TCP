package chap2

import (
	"fmt"
	"io"
	"log"
	"net"
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
	Listener, err := net.Listen("tcp", ":42069")
	if err != nil {
		log.Fatal(err)
	}
	defer Listener.Close()
	for {
		conn, err := Listener.Accept()
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		fmt.Println("connection accepted")
		for line := range getLinesChannel(conn) {
			fmt.Println(line)
		}
		fmt.Println("connection terminated")

	}

}
