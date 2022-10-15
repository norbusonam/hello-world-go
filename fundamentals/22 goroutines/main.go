package main

import (
	"fmt"
	"time"
)

func Count(num int) {
	for i := 1; i <= num; i++ {
		fmt.Println(i)
		time.Sleep(1000 * time.Millisecond)
	}
}

func main() {
	go Count(10)
	Count(3)
}
