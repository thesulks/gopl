package bank_test

import (
	"fmt"
	"os"
	"testing"

	bank "gopl/ch9/ex1"
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

	go func() {
		if ok := bank.Withdraw(300); !ok {
			fmt.Fprint(os.Stderr, "failed to withdraw\n")
		}
		done <- struct{}{}
	}()

	<-done

	if got, want := bank.Balance(), 0; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}

	if got, want := bank.Withdraw(1), false; got != want {
		t.Errorf("want to failed, but succeed in withdraw")
	}
}
