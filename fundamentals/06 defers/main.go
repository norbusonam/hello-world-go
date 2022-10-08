package main

import "fmt"

// defers build up in a stack, so they execute
// in reverse order of when they they occured
func countdown(start int) {
	fmt.Printf("Counting down from %d...\n", start)
	for i := 1; i <= start; i++ {
		defer fmt.Println(i)
	}
}

func main() {
	defer fmt.Println("main() is finished")
	countdown(3)
}
