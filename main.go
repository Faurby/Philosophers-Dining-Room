package main

import (
	"fmt"
	"time"
)

func main() {
	queryIN, queryOUT := make(chan string), make(chan string)

	ch1 := make(chan string)
	ch2 := make(chan string)
	ch3 := make(chan string)
	ch4 := make(chan string)
	ch5 := make(chan string)
	ch6 := make(chan string)
	ch7 := make(chan string)
	ch8 := make(chan string)
	ch9 := make(chan string)
	ch10 := make(chan string)
	ch11 := make(chan string)
	ch12 := make(chan string)
	ch13 := make(chan string)
	ch14 := make(chan string)
	ch15 := make(chan string)
	ch16 := make(chan string)
	ch17 := make(chan string)
	ch18 := make(chan string)
	ch19 := make(chan string)
	ch20 := make(chan string)

	f0 := Fork{id: 0, Lin: ch1, Lout: ch2, Rin: ch3, Rout: ch4, queryIN: queryOUT, queryOUT: queryIN}
	f1 := Fork{id: 1, Lin: ch5, Lout: ch6, Rin: ch7, Rout: ch8, queryIN: queryOUT, queryOUT: queryIN}
	f2 := Fork{id: 2, Lin: ch9, Lout: ch10, Rin: ch11, Rout: ch12, queryIN: queryOUT, queryOUT: queryIN}
	f3 := Fork{id: 3, Lin: ch13, Lout: ch14, Rin: ch15, Rout: ch16, queryIN: queryOUT, queryOUT: queryIN}
	f4 := Fork{id: 4, Lin: ch17, Lout: ch18, Rin: ch19, Rout: ch20, queryIN: queryOUT, queryOUT: queryIN}
	forkArray := []*Fork{&f0, &f1, &f2, &f3, &f4}

	pa := Philosopher{Name: "A", leftFork: &f0, rightFork: &f1, Lin: ch2, Lout: ch1, Rin: ch20, Rout: ch19, queryIN: queryOUT, queryOUT: queryIN}
	pb := Philosopher{Name: "B", leftFork: &f1, rightFork: &f2, Lin: ch6, Lout: ch5, Rin: ch4, Rout: ch3, queryIN: queryOUT, queryOUT: queryIN}
	pc := Philosopher{Name: "C", leftFork: &f2, rightFork: &f3, Lin: ch10, Lout: ch9, Rin: ch8, Rout: ch7, queryIN: queryOUT, queryOUT: queryIN}
	pd := Philosopher{Name: "D", leftFork: &f3, rightFork: &f4, Lin: ch14, Lout: ch13, Rin: ch12, Rout: ch11, queryIN: queryOUT, queryOUT: queryIN}
	pe := Philosopher{Name: "E", leftFork: &f4, rightFork: &f0, Lin: ch18, Lout: ch17, Rin: ch16, Rout: ch15, queryIN: queryOUT, queryOUT: queryIN}
	philArray := []*Philosopher{&pa, &pb, &pc, &pd, &pe}

	for _, element := range philArray {
		go element.start()
	}

	for _, element := range forkArray {
		go element.secondStart()
	}

	for {
		time.Sleep(time.Duration(1 * time.Second))
		for range philArray {
			queryOUT <- "timesEaten"
			input := <-queryIN

			if input != "go away" {
				fmt.Println(input)
			}
		}
		fmt.Println("----------------------------------")
		for range philArray {
			queryOUT <- "eating"
			input := <-queryIN

			if input != "go away" {
				fmt.Println(input)
			}
		}
		fmt.Println("----------------------------------")
		for range forkArray {
			queryOUT <- "timesUsed"
			input := <-queryIN

			if input != "go away" {
				fmt.Println(input)
			}
		}
		fmt.Println("----------------------------------")
	}
}
