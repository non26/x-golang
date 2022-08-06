package main

func main() {
	type Mock struct {
		x int
		y int
	}

	m := Mock{}
	if m == (Mock{}) {
		println("true")
	}

}
