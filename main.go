package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var lines []string

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		found := false
		for _, line := range lines {
			if line == text {
				found = true
				break
			}
		}
		if !found {
			lines = append(lines, text)
		}
	}
	if scanner.Err() != nil {
		fmt.Fprintf(os.Stderr, "%v: %v\n", os.Args[0], scanner.Err())
		os.Exit(1)
	}
	for _, line := range lines {
		fmt.Println(line)
	}
}
