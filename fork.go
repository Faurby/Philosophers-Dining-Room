package main

import "sync"

type Fork struct {
	id         int
	numberUsed int
	inUse      bool
	Lin        chan string
	Lout       chan string
	Rin        chan string
	Rout       chan string
}

func createFork(id int) {
	return Fork()
}
