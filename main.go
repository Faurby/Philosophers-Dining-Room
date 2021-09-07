package main

import "time"

func main() {
	//running := true
	//messageToPhil := make(chan string)
	//messageFromPhil := make(chan int)

	slice := MakePhils()

	//fmt.Println(slice)

	for _, element := range slice {
		go element.think()
	}

	time.Sleep(time.Duration(1000 * time.Second))
}

func MakePhils() []Philosopher {
	f1 := Fork{}
	f2 := Fork{}
	f3 := Fork{}
	f4 := Fork{}
	f5 := Fork{}

	p1 := Philosopher{Name: "Socrates", leftFork: &f1, rightFork: &f2}
	p2 := Philosopher{Name: "Sartre", leftFork: &f2, rightFork: &f3}
	p3 := Philosopher{Name: "Aristotle", leftFork: &f3, rightFork: &f4}
	p4 := Philosopher{Name: "Descartes", leftFork: &f4, rightFork: &f5}
	p5 := Philosopher{Name: "Pythagoras", leftFork: &f5, rightFork: &f1}
	return []Philosopher{p1, p2, p3, p4, p5}
}
