package main

import (
	"fmt"
	"unicode/utf8"
)

func reverse(b []byte) string {

	var rs []rune
	// 新しいメモリを割り当てないと[]runeを作れない
	// []runeがないと最後の文字のバイト数がわからない

	// 最後の文字をstringで取り出したとしても、最初の文字のバイト数<最後の文字だった場合に溢れる
	for len(b) > 0 {
		r, size := utf8.DecodeRune(b)
		// fmt.Printf("%c %v\n", r, size)
		rs = append(rs, r)

		b = b[size:]
	}

	for i, j := 0, len(rs)-1; i < j; i, j = i+1, j-1 {
		rs[i], rs[j] = rs[j], rs[i]
	}
	return string(rs)
}

func main() {
	a := []byte("あいうえおかき")
	fmt.Println(reverse(a[:]))
}
