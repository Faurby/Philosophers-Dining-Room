package main

import (
	"fmt"
	"sync"
)

func main() {
	var arbiter sync.Mutex
	running := true
	//messageToPhil := make(chan string)
	//messageFromPhil := make(chan int)

	slice := MakePhils()

	//fmt.Println(slice)

	for _, element := range slice {
		go element.eat()
	}

	for running {
		var consoleInput string
		fmt.Scan(&consoleInput)

		if consoleInput is not empty {

		}
	}
}
