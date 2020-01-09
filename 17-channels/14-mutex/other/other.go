package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

type accountType int

const (
	chequing accountType = iota
	saving
	investing
	current
	rrsp
)

func (at accountType) String() string {
	switch at {
	case chequing:
		return "chequing"
	case saving:
		return "saving"
	case investing:
		return "investing"
	case current:
		return "current"
	case rrsp:
		return "rrsp"
	default:
		return "saving"
	}
}

type account struct {
	number  string
	purpose accountType
	balance float64
}

func (a account) String() string {
	s := fmt.Sprintf("%s:%s %.2f", a.purpose, a.number, a.balance)
	return s
}
func useAtm(a *account) {
	val := float64(randInt(999)) + float64(randInt(100))/100.00

	transaction := rand.Intn(9)
	if transaction > 4 {
		a.balance -= val
	} else {
		a.balance += val
	}
}

var sd = rand.NewSource(time.Now().UnixNano())

//UnsignedInt generates a random unsigned int up to max value
func randInt(max int) int {
	r := rand.New(sd)
	n := r.Intn(max)
	return n
}
func newAccount() account {
	number := fmt.Sprintf("%d%d%d-%d%d%d%d-%d", randInt(9), randInt(9), randInt(9), randInt(9), randInt(9), randInt(9), randInt(9), randInt(9))
	purpose := accountType(randInt(int(rrsp) + 1))
	balance := float64(randInt(9999)) + float64(randInt(100))/100.00
	return account{number: number, purpose: purpose, balance: balance}
}
func main() {

	// For our example the `accounts` will be a map.
	var accounts = make(map[string]account)
	var accountNumbers []string
	for n := 0; n < 6; n++ {
		a := newAccount()
		number := a.number
		accountNumbers = append(accountNumbers, number)
		accounts[number] = a
	}
	// This `mutex` will synchronize access to `accounts`.
	var mutex = &sync.Mutex{}

	// We'll keep track of how many read and write
	// operations we do.
	var reads uint64
	var writes uint64

	// Here we start 100 goroutines to execute repeated
	// reads against the accounts, once per millisecond in
	// each goroutine.
	for r := 0; r < 100; r++ {
		go func() {
			total := 0.0
			for {

				// For each read we pick a key to access,
				// `Lock()` the `mutex` to ensure
				// exclusive access to the `accounts`, read
				// the value at the chosen key,
				// `Unlock()` the mutex, and increment
				// the `reads` count.
				ix := rand.Intn(5)
				key := accountNumbers[ix]
				mutex.Lock()
				total += accounts[key].balance
				mutex.Unlock()
				atomic.AddUint64(&reads, 1)

				// Wait a bit between reads.
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// We'll also start 10 goroutines to simulate writes,
	// using the same pattern we did for reads.
	for w := 0; w < 10; w++ {
		go func() {
			for {
				ix := rand.Intn(5)
				key := accountNumbers[ix]
				mutex.Lock()
				a := accounts[key]
				useAtm(&a)
				accounts[key] = a
				mutex.Unlock()
				atomic.AddUint64(&writes, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// Let the 10 goroutines work on the `accounts` and
	// `mutex` for a second.
	time.Sleep(time.Second)

	// Take and report final operation counts.
	readsFinal := atomic.LoadUint64(&reads)
	fmt.Println("balance checks:", readsFinal)
	writesFinal := atomic.LoadUint64(&writes)
	fmt.Println("transactions:", writesFinal)

	// With a final lock of `accounts`, show how it ended up.
	mutex.Lock()
	for account := range accounts {
		fmt.Println("accounts:", accounts[account])
	}
	mutex.Unlock()
}
