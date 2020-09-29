package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func getUnicodeCategory(r rune) rune {
	for name, table := range unicode.Categories {
		if len(name) == 2 && unicode.Is(table, r) {
			return []rune(name)[0]
		}
	}
	return rune('C')
}

func main() {
	counts := make(map[rune]int)
	var utflen [utf8.UTFMax + 1]int
	categories := make(map[rune]int)
	invalid := 0

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
		categories[getUnicodeCategory(r)]++
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	fmt.Printf("\ncategory\tcount\n")
	for c, n := range categories {
		fmt.Printf("%s\t%d\n", string(c), n)
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
