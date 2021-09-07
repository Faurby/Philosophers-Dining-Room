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
}

func (p Philosopher) eat() {
	p.eating = true
	//fmt.Println(p.Name + " is eating")
	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	p.timesEaten++
	//fmt.Println(p.Name + " has eaten " + strconv.Itoa(p.timesEaten) + " times")
	p.eating = false
	p.think()
}

func (p Philosopher) think() {
	p.thinking = true
	//fmt.Println(p.Name + " is thinking")
	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	p.thinking = false
	p.eat()
}

func MakePhils() []Philosopher {
	p1 := Philosopher{Name: "Socrates"}
	p2 := Philosopher{Name: "Sartre"}
	p3 := Philosopher{Name: "Aristotle"}
	p4 := Philosopher{Name: "Descartes"}
	p5 := Philosopher{Name: "Pythagoras"}
	return []Philosopher{p1, p2, p3, p4, p5}
}
