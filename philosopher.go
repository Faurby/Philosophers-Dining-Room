package main

import (
	"fmt"
	"time"
)

type Philosopher struct {
	Name       string
	timesEaten int
	rightFork  *Fork
	leftFork   *Fork
	Lin        chan string
	Lout       chan string
	Rin        chan string
	Rout       chan string
}

func (p *Philosopher) start() {
	for {
		if p.Name == "A" || p.Name == "C" {
			p.Rout <- "i want to eat"
			<-p.Rin
			p.Lout <- "i want to eat"
			<-p.Lin
			fmt.Println(p.Name + " eating")
			time.Sleep(time.Duration(3 * time.Second))
			p.timesEaten++
			p.Rout <- "im done"
			p.Lout <- "im done"
			fmt.Println(p.Name + " thinking")
			time.Sleep(time.Duration(3 * time.Second))

		} else {
			p.Lout <- "i want to eat"
			<-p.Lin
			p.Rout <- "i want to eat"
			<-p.Rin
			fmt.Println(p.Name + " eating")
			time.Sleep(time.Duration(3 * time.Second))
			p.timesEaten++
			p.Lout <- "im done"
			p.Rout <- "im done"
			fmt.Println(p.Name + " thinking")
			time.Sleep(time.Duration(3 * time.Second))
		}
	}
}
