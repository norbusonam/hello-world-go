package main

import "fmt"

func HowMany[T comparable](s []T, x T) int {
	count := 0
	for _, val := range s {
		if val == x {
			count++
		}
	}
	return count
}

func main() {

	// strings
	names := []string{"Jack", "Alex", "Ted", "Alex"}
	numOfJacks := HowMany(names, "Jack")
	numOfAlexs := HowMany(names, "Alex")
	numOfTeds := HowMany(names, "Ted")
	fmt.Println(names)
	fmt.Printf("Number of people named %s: %d\n", "Jack", numOfJacks)
	fmt.Printf("Number of people named %s: %d\n", "Alex", numOfAlexs)
	fmt.Printf("Number of people named %s: %d\n", "Ted", numOfTeds)

	// numbers
	numbers := []int{12, 14, 12, 24, 12, 14}
	numOf12s := HowMany(numbers, 12)
	numOf14s := HowMany(numbers, 14)
	numOf24s := HowMany(numbers, 24)
	fmt.Println(numbers)
	fmt.Printf("Number of %d's: %d\n", 12, numOf12s)
	fmt.Printf("Number of %d's: %d\n", 14, numOf14s)
	fmt.Printf("Number of %d's: %d\n", 24, numOf24s)

}
