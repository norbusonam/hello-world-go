package main

import "fmt"

type Any interface{}

func describe(val Any) {
	fmt.Printf("%v has type %T\n", val, val)
}

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

	var pi Any = 3.14159
	describe(pi)

	anotherPi := pi.(float64)
	describe(anotherPi)

	// the commented out line below will trigger a
	// panic, the second return value must be accepted
	// to avoid it
	// s := pi.(string)
	s, ok := pi.(string)
	describe(s)
	describe(ok)

}
