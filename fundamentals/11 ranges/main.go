package main

import "fmt"

func main() {

	primes := []int{2, 3, 5, 7, 11, 13}

	// index and value
	for i, prime := range primes {
		fmt.Printf("%d is at index %d\n", prime, i)
	}

	// only value
	for _, prime := range primes {
		fmt.Printf("%d is a prime\n", prime)
	}

	// only index
	for i := range primes {
		fmt.Printf("index %d\n", i)
	}

}
