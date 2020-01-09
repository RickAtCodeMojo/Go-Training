package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var seed = rand.NewSource(time.Now().UnixNano())

//UnsignedInt generates a random unsigned int up to max value
func number(max int) int {
	r1 := rand.New(seed)
	n := r1.Intn(max)
	return n
}

func sleepRandom(milliseconds int, random int) {
	time.Sleep(time.Millisecond * time.Duration(milliseconds*number(random)))

}

type monsterKind int

const (
	Orc monsterKind = iota
	Troll
	Baselisk
	Zombie
)

type monster struct {
	name   string
	isA    monsterKind
	health int
	damage int
}

// Here’s the worker, of which we’ll run several concurrent instances.
// These workers will receive work on the jobs channel and send the corresponding
// results on results. We’ll sleep a second per job to simulate an expensive task.
func worker(wg *sync.WaitGroup, id int, A <-chan monster, B <-chan monster, survivors chan<- monster) {
	defer wg.Done()
	for a := range A {
		for b := range B {
			fmt.Println("worker", id, "Fight:", a.name, "vs", b.name, "!")
			a.health -= b.damage
			b.health -= a.damage
			fmt.Println("worker", id, "Results- A:", a.name, ", health:", a.health, ", B:", b.name, ", health:", b.health)
			if a.health > b.health {
				survivors <- a
			} else {
				survivors <- b
			}
		}
	}

}
func main() {

	teamA := make(chan monster, 100)
	teamB := make(chan monster, 100)
	// In order to use our pool of workers we need to send them work and collect
	//their results. We make 2 channels for this.
	//jobs := make(chan int, 100)
	survivors := make(chan monster, 100)

	// This starts up 3 workers, initially blocked because there are no jobs yet.
	var wg sync.WaitGroup
	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go worker(&wg, w, teamA, teamB, survivors)
	}

	// Here we send 5 jobs and then close that channel to indicate that’s all
	//the work we have.
	for j := 1; j <= 5; j++ {
		teamA <- monster{name: fmt.Sprintf("A%d", j), isA: monsterKind(number(int(Zombie))), health: number(50), damage: number(10)}
		teamB <- monster{name: fmt.Sprintf("B%d", j), isA: monsterKind(number(int(Zombie))), health: number(50), damage: number(10)}
	}
	close(teamA)
	close(teamB)
	wg.Wait()
	for j := 1; j <= 5; j++ {
		fmt.Println(<-survivors)
	}
}
