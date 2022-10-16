package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(time.Second)
	boom := time.After(time.Second * 3)
	for {
		select {
		case <-tick:
			fmt.Printf("tick")
		case <-boom:
			fmt.Println("\nBOOM")
			return
		default:
			fmt.Printf(".")
			time.Sleep(time.Millisecond * 100)
		}
	}
}
