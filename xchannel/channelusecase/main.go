package main

// https://go101.org/article/channel-use-cases.html
// https://stackoverflow.com/questions/31920353/whats-the-difference-between-chan-and-chan-as-a-function-return-type

import (
	"crypto/rand"
	"fmt"
	"os"
	"sort"
)

func main() {
	values := make([]byte, 32*1024*1024)
	if _, err := rand.Read(values); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	done := make(chan struct{}) // can be buffered or not

	// The sorting goroutine
	go func() {
		sort.Slice(values, func(i, j int) bool {
			return values[i] < values[j]
		})
		// Notify sorting is done.
		done <- struct{}{}
	}()

	// do some other things ...

	<-done // waiting here for notification
	fmt.Println(values[0], values[len(values)-1])
}
