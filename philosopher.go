package main

import (
	"math/rand"
	"time"
)

func channel() chan int {
	return make(chan int)
}

func checkChannel() int {
	var x int
	c := channel()
	x = <-c
	if x != 0 {
		return x
	} else {
		return -1
	}
}

type Philosopher struct {
	Name       string
	timesEaten int
	eating     bool
	thinking   bool
	rightFork  *Fork
	leftFork   *Fork
}

func (p *Philosopher) eat() {
	//check if I am eating
	for !p.eating {
		//check if both forks are available
		if !p.rightFork.inUse && !p.leftFork.inUse {
			p.rightFork.Lock()
			p.leftFork.Lock()

			//eating is true, go out of while loop next time
			p.eating = true
			//fmt.Println(p.Name + " is eating")
			time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
			p.timesEaten++
			//fmt.Println(p.Name + " has eaten " + strconv.Itoa(p.timesEaten) + " times")

			p.rightFork.Unlock()
			p.leftFork.Unlock()
		}
	}
	//eating is false, now think
	p.eating = false
	p.think()
}

func (p *Philosopher) think() {
	p.thinking = true
	//fmt.Println(p.Name + " is thinking")
	time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
	p.thinking = false
	p.eat()
}
