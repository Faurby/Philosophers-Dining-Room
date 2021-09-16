package main

import (
	"fmt"
	"strconv"
	"time"
)

type Philosopher struct {
	Name       string
	timesEaten int
	rightFork  *Fork
	leftFork   *Fork
	eating     bool
	Lin        chan string
	Lout       chan string
	Rin        chan string
	Rout       chan string
	queryIN    chan string
	queryOUT   chan string
}

func (p *Philosopher) handleQuery(in chan string, out chan string) {
	for {
		x := <-in
		if x == "timesEaten" {
			var s = fmt.Sprintf("Philosopher %s has eaten %s times", p.Name, strconv.Itoa(p.timesEaten))

			out <- s
		} else if x == "eating" {
			if p.eating {
				out <- fmt.Sprintf("%s is eating", p.Name)
			} else {
				out <- fmt.Sprintf("%s is thinking", p.Name)
			}
		} else {
			out <- "go away"
		}
	}

}

func (p *Philosopher) start() {
	go p.handleQuery(p.queryIN, p.queryOUT)
	for {
		if p.Name == "A" || p.Name == "C" {
			p.Rout <- "i want to eat"
			<-p.Rin
			p.Lout <- "i want to eat"
			<-p.Lin
			//fmt.Println(p.Name + " eating")
			p.eating = true
			delay()
			p.timesEaten++
			p.Rout <- "im done"
			p.Lout <- "im done"
			//fmt.Println(p.Name + " thinking")
			p.eating = false
			delay()
		} else {
			p.Lout <- "i want to eat"
			<-p.Lin
			p.Rout <- "i want to eat"
			<-p.Rin
			//fmt.Println(p.Name + " eating")
			p.eating = true
			delay()
			p.timesEaten++
			p.Lout <- "im done"
			p.Rout <- "im done"
			p.eating = false
			//fmt.Println(p.Name + " thinking")
			delay()
		}
	}
}

func delay() {
	time.Sleep(time.Duration(3 * time.Second))
}
