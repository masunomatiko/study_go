package bank_test

import (
	"fmt"
	"testing"

	bank "github.com/masunomatiko/study_go/ch09/ex1"
)

func TestBank(t *testing.T) {
	done := make(chan struct{})

	// Alice
	go func() {
		bank.Deposit(200)
		fmt.Println("=", bank.Balance())
		done <- struct{}{}
	}()

	// Bob
	go func() {
		bank.Deposit(100)
		done <- struct{}{}
	}()

	// Wait for both transactions.
	<-done
	<-done

	// withdraw test
	// Lana
	go func() {
		fmt.Println(bank.WithDraw(150))
		done <- struct{}{}
	}()
	<-done

	// Tom
	go func() {
		fmt.Println(bank.WithDraw(200))
		done <- struct{}{}
	}()
	<-done

	if got, want := bank.Balance(), 150; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
}
