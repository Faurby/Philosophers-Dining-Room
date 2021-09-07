package main

import "sync"

func main() {
	var arbiter sync.Mutex

	arbiter.Lock()

	arbiter.Unlock()
}
