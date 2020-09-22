package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func commaWithFloat(s string) string {
	var buf bytes.Buffer
	dIdx := 0
	if string(s[0]) == "+" || string(s[0]) == "-" {
		buf.WriteByte(s[0])
		dIdx = 1
	}
	pIdx := strings.Index(s, ".")
	idx := (pIdx - dIdx) % 3
	if idx == 0 {
		idx = 3
	}
	buf.WriteString(s[dIdx : dIdx+idx])
	for i := dIdx + idx; i < pIdx; i += 3 {
		buf.WriteByte(',')
		buf.WriteString(string(s[i : i+3]))
	}
	buf.WriteString(s[pIdx:])
	return buf.String()
}

func comma(s string) string {
	var buf bytes.Buffer

	return buf.String()
}

func main() {
	text := os.Args[1]
	fmt.Printf("%s\n", commaWithFloat(text))
}
