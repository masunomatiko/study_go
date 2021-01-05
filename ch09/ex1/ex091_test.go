package bank

import (
	"fmt"
	"testing"
)

func TestBank(t *testing.T) {
	done := make(chan struct{})

	go func() {
		Deposit(200)
		done <- struct{}{}
	}()

	go func() {
		Deposit(100)
		done <- struct{}{}
	}()

	<-done
	<-done

	go func() {
		fmt.Println(WithDraw(150))
		done <- struct{}{}
	}()
	<-done

	// Withdrawal failure
	go func() {
		fmt.Println(WithDraw(200))
		done <- struct{}{}
	}()
	<-done
	balance := Balance()
	expected := 150
	if balance != expected {
		t.Errorf("Balance = %d, expected %d", balance, expected)
	}
}
