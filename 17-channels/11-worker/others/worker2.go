package main

import (
	"fmt"
	"math/rand"
	"time"
)

type activities int

const (
	deposit activities = iota
	withdrawal
)

type transaction struct {
	account  string
	activity activities
	amount   float64
}

type account struct {
	number  string
	balance float64
}

var accounts []account
var sd = rand.NewSource(time.Now().UnixNano())

//UnsignedInt generates a random unsigned int up to max value
func randInt(max int) int {
	r := rand.New(sd)
	n := r.Intn(max)
	return n
}

//new generates a new account
func newAccount() account {
	number := fmt.Sprintf("%d%d%d-%d%d%d%d-%d", randInt(9), randInt(9), randInt(9), randInt(9), randInt(9), randInt(9), randInt(9), randInt(9))
	balance := 0.00
	return account{number: number, balance: balance}
}

func amount() float64 {
	a := float64(randInt(2000))
	change := float64(randInt(100))
	if change > 0 {
		a += 100.0 / change
	}
	return a
}
func activity() activities {
	a := randInt(99)
	if a > 49 {
		return deposit
	} else {
		return withdrawal
	}
}
func (a activities) String() string {
	switch a {
	case deposit:
		return "Deposit"
	case withdrawal:
		return "Withdrawal"
	default:
		return "UNKNOWN"
	}
}

var runningBalance float64

const depositFee = 1.50
const withdrawalFee = 2.50

// Here’s the worker, of which we’ll run several concurrent instances.
// These workers will receive work on the jobs channel and send the corresponding
// results on results. We’ll sleep a second per job to simulate an expensive task.
func useAtm(id int, A <-chan transaction, balance chan<- transaction) {
	for a := range A {
		var fee float64
		time.Sleep(time.Second)
		switch a.activity {
		case deposit:
			fee = depositFee
			a.amount -= depositFee
		case withdrawal:
			fee = withdrawalFee
			a.amount += withdrawalFee
		default:
		}
		balance <- a
		fmt.Printf("ATM:%d. %s on account:%s for %.2f (includes fee:%.2f),\n", id, a.activity, a.account, a.amount, fee)
	}

}
func reconciler(a *account, bal chan transaction, done chan bool) {
	for j := 1; j <= 5; j++ {
		t := <-bal
		if t.activity == withdrawal {
			a.balance -= t.amount
		} else {
			a.balance += t.amount
		}
		fmt.Printf("Running Balance:%.2f\n", a.balance)
	}
}

func main() {
	runningBalance = 100
	done := make(chan bool)
	//make 5 accounts
	for a := 0; a < 5; a++ {
		accounts = append(accounts, newAccount())
	}
	// In order to use our pool of workers we need to send them work and collect
	//their results. We make 2 channels for this.
	transactions := make(chan transaction, 100)
	balance := make(chan transaction, 100)

	// This starts up 3 workers, initially blocked because there are no jobs yet.
	for w := 1; w <= 3; w++ {
		go useAtm(w, transactions, balance)
	}

	// Here we send 5 jobs and then close that channel to indicate that’s all
	//the work we have.
	checking := accounts[0]
	for j := 1; j <= 5; j++ {
		transactions <- transaction{account: checking.number, amount: amount(), activity: activity()}
	}

	reconciler(&checking, balance, done)

	// close(transactions)
	fmt.Printf("Closing Balance:%.2f\n", checking.balance)
}
