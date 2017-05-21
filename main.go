package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
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
	} else {
		b, err := ioutil.ReadFile(os.Args[1])
		if err != nil {
			return 1
		}
		fmt.Print(string(b))
	}
	return 0
}
