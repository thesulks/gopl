// Package bank provides a concurrency-safe bank with one account.
package bank

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance
var withdraws = make(chan withdrawMessage)

type withdrawMessage struct {
	ok     chan bool
	amount int
}

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }
func Withdraw(amount int) bool {
	ok := make(chan bool)
	withdraws <- withdrawMessage{ok, amount}
	return <-ok
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case withdraw := <-withdraws:
			if balance < withdraw.amount {
				withdraw.ok <- false
				continue
			}
			balance -= withdraw.amount
			withdraw.ok <- true
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}
