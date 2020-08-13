package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	counts := make(map[string][]string)
	// filenames := make(map[string][]string)
	files := os.Args[1:]

	for _, arg := range files {
		f, err := os.Open(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			continue
		}
		countLines(f, counts, arg)
		f.Close()
	}
	for line, filenames := range counts {
		if len(filenames) > 1 {
			fmt.Printf("%s\t%s\n", line, filenames)
		}
	}
}

func countLines(f *os.File, counts map[string][]string, arg string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()] = append(counts[input.Text()], filepath.Base(arg))
	}
}
