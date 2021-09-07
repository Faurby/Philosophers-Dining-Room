package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Philosopher struct {
	Name       string
	timesEaten int
	eating     bool
	thinking   bool
}

func (p Philosopher) eat() {
	p.eating = true
	fmt.Println(p.Name + " is eating")
	time.Sleep(time.Duration(rand.Intn(5)) * time.Millisecond)
	p.timesEaten++
	fmt.Println(p.Name + " has eaten " + strconv.Itoa(p.timesEaten) + " times")
	p.eating = false
	p.think()
}

func (p Philosopher) think() {
	p.thinking = true
	fmt.Println(p.Name + " is thinking")
	time.Sleep(time.Duration(rand.Intn(5)) * time.Millisecond)
	p.thinking = false
	p.eat()
}

func makePhils() []Philosopher {
	p1 := Philosopher{Name: "Socrates"}
	p2 := Philosopher{Name: "Sartre"}
	p3 := Philosopher{Name: "Aristotle"}
	p4 := Philosopher{Name: "Descartes"}
	p5 := Philosopher{Name: "Pythagoras"}
	return []Philosopher{p1, p2, p3, p4, p5}
}
