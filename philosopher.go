package main

import (
	"fmt"
	"math/rand"
	"strconv"
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
	rightFork  Fork
	leftFork   Fork
}

func (p Philosopher) eat() {
	//check if I can eat
	if !p.rightFork.inUse && !p.leftFork.inUse {
		p.rightFork.arbiter.Lock()
		p.leftFork.arbiter.Lock()

		p.eating = true
		fmt.Println(p.Name + " is eating")
		time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
		p.timesEaten++
		fmt.Println(p.Name + " has eaten " + strconv.Itoa(p.timesEaten) + " times")
		p.eating = false
		p.rightFork.arbiter.Unlock()
		p.leftFork.arbiter.Unlock()
		p.think()
	}
}

func (p Philosopher) think() {
	p.thinking = true
	fmt.Println(p.Name + " is thinking")
	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	p.thinking = false
	p.eat()
}
