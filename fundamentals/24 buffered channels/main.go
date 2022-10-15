package main

import "fmt"

func main() {

	c := make(chan int, 2)
	c <- 1
	c <- 2
	// the following line would cause failure
	// c <- 3
	fmt.Println(<-c)
	fmt.Println(<-c)

}
