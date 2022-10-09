package main

import "fmt"

func main() {

	// arrays are fixed length
	var primes [6]int = [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// arrays are value types
	primes2 := primes
	primes2[5] = 17
	fmt.Println(primes)
	fmt.Println(primes2)

}
