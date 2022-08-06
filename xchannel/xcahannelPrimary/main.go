package main

import (
	"fmt"
	"time"
)

func worker(ch chan int, chQuit chan struct{}) {
	println("enter")
	for {
		select {
		case v := <-ch:
			fmt.Printf("Got value %d\n", v)
		case <-chQuit:
			fmt.Printf("Signalled on quit channel. Finishing\n")
			time.Sleep(5 * time.Second)
			chQuit <- struct{}{}
			return
		}
		println("loop")
	}
}
func main() {
	ch, chQuit := make(chan int), make(chan struct{})
	go worker(ch, chQuit)
	ch <- 3
	chQuit <- struct{}{}

	// wait to be signalled back by the worker
	// the code below this will wait until there's value to chQuit channel
	<-chQuit
	println("after")
}
