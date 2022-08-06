package main

import "fmt"

func gen(nums []int) <-chan struct{} {
	out := make(chan struct{})
	go func() {
		// for _, n := range nums {
		// 	out <- n
		// }
		close(out)
	}()
	fmt.Println("return statement is called ")
	return out
}

func main() {
	// c := make(chan int)
	c := gen([]int{2, 3, 4, 5})
	_ = c

	// Consume the output.
	// Print 2,3,4,5
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	println("is blocked")

}
