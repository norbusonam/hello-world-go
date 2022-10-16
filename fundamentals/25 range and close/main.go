package main

import "fmt"

func fib(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	// close to indicate no more values
	// will be sent
	close(c)
}

func main() {
	c := make(chan int, 20)
	fib(cap(c), c)
	// loop until channel closed
	for num := range c {
		fmt.Println(num)
	}
}
