package main

import "fmt"

type Any interface{}

func sayType(something Any) {
	switch something.(type) {
	case int:
		fmt.Println("It's an int!")
	case float32:
		fmt.Println("It's an float32!")
	case float64:
		fmt.Println("It's an float64!")
	case string:
		fmt.Println("It's an string!")
	default:
		fmt.Println("I don't know what it is.")
	}
}

func main() {
	sayType("cool string bro")
	sayType(22)
	sayType(2.718)
	sayType([]int{2, 4, 6, 8})
}
