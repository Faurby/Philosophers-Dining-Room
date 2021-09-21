package main

import (
	"fmt"
	"strconv"
)

type Fork struct {
	id       int
	used     int
	Lin      chan string
	Lout     chan string
	Rin      chan string
	Rout     chan string
	queryIN  chan string
	queryOUT chan string
}

func (F *Fork) handleQuery(in chan string, out chan string) {
	for {
		x := <-in
		if x == "timesUsed" {
			var s = fmt.Sprintf("Fork %s has been used %s times", strconv.Itoa(F.id), strconv.Itoa(F.used))
			out <- s
		} else {
			out <- "go away"
		}
	}
}

func (F *Fork) secondStart() {
	go F.handleQuery(F.queryIN, F.queryOUT)
	for {
		select {
		case <-F.Lin:
			F.Lout <- "eat now"
			<-F.Lin
			F.used++
		case <-F.Rin:
			F.Rout <- "eat now"
			<-F.Rin
			F.used++
		}
	}
}
