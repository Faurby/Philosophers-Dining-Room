package main

import "sync"

type Fork struct {
	numberUsed int
	inUse      bool
	sync.Mutex
}
