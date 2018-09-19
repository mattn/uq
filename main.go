package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if !isStdin() {
		fmt.Fprintln(os.Stderr, "stdin not detected")
		os.Exit(1)
	}

	lines := make(map[string]struct{})

	scanner := bufio.NewScanner(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	var text string
	var ok bool
	for scanner.Scan() {
		text = scanner.Text()
		if _, ok = lines[text]; !ok {
			lines[text] = struct{}{}
			out.WriteString(text + "\n")
		}
	}
	if scanner.Err() != nil {
		fmt.Fprintf(os.Stderr, "%v: %v\n", os.Args[0], scanner.Err())
		os.Exit(1)
	}
	out.Flush()
}

func isStdin() bool {
	stat, err := os.Stdin.Stat()
	if err != nil {
		return false
	}
	return (stat.Mode() & os.ModeCharDevice) == 0
}
