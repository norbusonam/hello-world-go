package main

import "fmt"

// basic
func add(x int, y int) int {
	return x + y
}

// simplified params
func subtract(x, y int) int {
	return x - y
}

// simplified params and named return value
func multiply(x, y int) (product int) {
	product = x * y
	return
}

func main() {
	var x int = 6
	var y int = 3
	fmt.Println(add(x, y))
	fmt.Println(subtract(x, y))
	fmt.Println(multiply(x, y))
}
