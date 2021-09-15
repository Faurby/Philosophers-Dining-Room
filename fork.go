package main

type Fork struct {
	id   int
	Lin  chan string
	Lout chan string
	Rin  chan string
	Rout chan string
}

func (F *Fork) secondStart() {
	for {
		select {
		case <-F.Lin:
			F.Lout <- "eat now"
			<-F.Lin
		case <-F.Rin:
			F.Rout <- "eat now"
			<-F.Rin
		}
	}
}
