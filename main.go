package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	lines := make(map[string]struct{})

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		if _, ok := lines[text]; !ok {
			lines[text] = struct{}{}
			fmt.Println(text)
		}
	}
	if scanner.Err() != nil {
		fmt.Fprintf(os.Stderr, "%v: %v\n", os.Args[0], scanner.Err())
		os.Exit(1)
	}
}
