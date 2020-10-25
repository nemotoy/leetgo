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
		line, _, err := reader.ReadRune()
		if err == io.EOF {
			break
		}
		if line == '\n' || line == ',' {
			continue
		}
		fmt.Println("#", string(line))
	}
}
