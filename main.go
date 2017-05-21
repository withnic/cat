package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	os.Exit(run())
}

func run() int {
	if len(os.Args) == 1 {
		reader := bufio.NewReader(os.Stdin)
		for {
			text, err := reader.ReadString('\n')
			if err != nil {
				return 1
			}
			fmt.Print(text)
		}
	}
	return 0
}
