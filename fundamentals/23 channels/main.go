package main

import (
	"fmt"
	"time"
)

func sum(nums []int, c chan int) {
	sum := 0
	for _, num := range nums {
		sum += num
		time.Sleep(time.Second)
	}
	c <- sum
}

func main() {
	c := make(chan int)

	// run goroutines
	go sum([]int{1, 1, 1, 1}, c)
	go sum([]int{1, 1}, c)

	// block until channel receives two values
	x, y := <-c, <-c
	fmt.Println(x)
	fmt.Println(y)
}
