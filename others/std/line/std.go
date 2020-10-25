package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {

	// const data = "aaa\nbbb\nccc"
	// in := strings.NewReader(data)

	reader := bufio.NewReader(os.Stdin)

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		fmt.Println("#", string(line))
	}
}
