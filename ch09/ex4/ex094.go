package main

import (
	"fmt"
	"time"
)

// 限界値みる方法わからなかった
const stageLimit = 1000000

// yにxをどんどんつなげていくイメージ
func pipeline(stage int) (x chan int, y chan int) {
	// create channel
	y = make(chan int)
	first := y
	for i := 0; i < stage; i++ {
		x = y
		// create channel
		y = make(chan int)
		go func(x chan int, y chan int) {
			for v := range x {
				// add
				y <- v
			}
			close(y)
		}(x, y)
	}
	return first, y

}

func main() {
	x, y := pipeline(stageLimit)
	s := time.Now()
	x <- 1
	<-y
	close(x)
	fmt.Printf("%d goroutines, %fs\n", stageLimit, time.Since(s).Seconds())
}
