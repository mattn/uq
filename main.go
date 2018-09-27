package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func run(r io.Reader, w io.Writer) int {
	lines := make(map[string]struct{})

	scanner := bufio.NewScanner(r)
	out := bufio.NewWriter(w)
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
		return 1
	}
	out.Flush()
	return 0
}

func isStdin() bool {
	stat, err := os.Stdin.Stat()
	if err != nil {
		return false
	}
	return (stat.Mode() & os.ModeCharDevice) == 0
}

func main() {
	if !isStdin() {
		fmt.Fprintln(os.Stderr, "stdin not detected")
		os.Exit(1)
	}
	os.Exit(run(os.Stdin, os.Stdout))
}
