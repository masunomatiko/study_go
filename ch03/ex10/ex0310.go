package main

import (
	"bytes"
	"fmt"
	"os"
)

func comma(s string) string {
	var buf bytes.Buffer
	idx := len(s) % 3
	if idx == 0 {
		idx = 3
	}
	buf.WriteString(s[:idx])
	for i := idx; i < len(s); i += 3 {
		buf.WriteByte(',')
		buf.WriteString(string(s[i : i+3]))
	}

	return buf.String()
}

func main() {
	text := os.Args[1]
	fmt.Printf("%s\n", comma(text))
}
