package main

import (
	"fmt"
	"time"
)

func main() {

	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			default:
				// println("default")
			}
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	println(1)
	ticker.Stop()
	println(2)
	// done <- true
	// <-done //// error
	fmt.Println("Ticker stopped")
}
