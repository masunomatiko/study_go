package main

import "fmt"

func rotate(a []int) {
	tmp := a[0]
	copy(a, a[1:])
	a[len(a)-1] = tmp
}

func main() {
	s := []int{0, 1, 2, 3, 4, 5}
	rotate(s)
	fmt.Println(s)
}
