package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func getLinesChannel(f io.ReadCloser) <-chan string {
	// channel with buffer size of 1
	out := make(chan string, 1)
	
	// goroutines
	go func() {
		defer f.Close()
		defer close(out)

		str := ""
		for {
		data := make([]byte,8)
		n, err := f.Read(data)
		if err != nil {
			break
		}
		
		data = data[:n]
		if i := bytes.IndexByte(data, '\n'); i != -1 {
			str += string(data[:i])
			data = data[i+1:]
			out <- str
			str = ""
		}
		str += string(data)
	}

	if len(str) != 0 {
		out <- str
	}
	}()

	return out
}


func main() {
	file, err := os.Open("message.txt")
	if err != nil {
		log.Println("Error reading file: ", err)
	}
	
	lines := getLinesChannel(file)
	for line:= range lines {
		fmt.Printf("read: %s\n", line)
	}
	
	
}
