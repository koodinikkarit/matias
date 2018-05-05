package main

import "fmt"

type Tyyppi struct {
	C chan int
}

type Channels struct {
	kanava1 chan int
	kanava2 chan int
}

func main() {
	channels := Channels{
		kanava1: make(chan int),
		kanava2: make(chan int),
	}

	go func() {
		for {
			select {
			case i := <-channels.kanava1:
				fmt.Println("arvo ", i)
			}
		}
	}()

	go func() {
		for {
			select {
			case i := <-channels.kanava2:
				fmt.Println("value ", i)
			}
		}
	}()

	//channels.kanava1 = make(chan int)
	channels.kanava1 <- 50000
	channels.kanava2 <- 50000

	for {

	}
}
