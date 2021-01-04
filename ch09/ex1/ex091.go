package main

import (
	"fmt"
)

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance
var withdraw = make(chan int) // send amount to withdraw
var result = make(chan bool)  // withdraw result

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }
func WithDraw(amount int) bool {
	withdraw <- amount
	return <-result
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case amount := <-withdraw:
			if amount <= balance {
				balance -= amount
				result <- true
			} else {
				result <- false
			}
		}
	}
}

// func init() {

// }

func main() {
	go teller() // start the monitor goroutine
	done := make(chan struct{})

	// Alice
	go func() {
		Deposit(200)
		done <- struct{}{}
	}()

	// Bob
	go func() {
		Deposit(100)
		done <- struct{}{}
	}()

	// Wait for both transactions.
	<-done
	<-done

	// withdraw test
	// Lana
	go func() {
		fmt.Println(WithDraw(150))
		done <- struct{}{}
	}()
	<-done

	// Tom
	go func() {
		fmt.Println(WithDraw(200))
		done <- struct{}{}
	}()
	<-done
	balance := Balance()
	fmt.Println(balance)
}
