package main

import (
	"fmt"
	"unicode"
)

func remove_dup_spaces(a []string) []string {
	var idx int
	for _, s := range a {
		if unicode.IsSpace([]rune(s)[0]) && a[idx] == " " {
			continue
		}
		if unicode.IsSpace([]rune(s)[0]) && s != " " {
			a[idx] = " "
		} else {
			a[idx] = s
		}
		idx++

		fmt.Println(a[:idx])
	}
	return a[:idx]
}

func main() {
	a := []string{"a", "b", "\n", "\n", "c", " "}
	a = remove_dup_spaces(a)
	fmt.Println(a)
}
