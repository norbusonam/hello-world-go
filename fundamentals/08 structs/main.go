package main

import "fmt"

type Point struct {
	Name string
	X    int
	Y    int
}

func printPoint(point Point) {
	fmt.Printf("%s is located at point %d, %d\n", point.Name, point.X, point.Y)
}

func printPointerToPoint(pointerToPoint *Point) {
	fmt.Printf("%s is located at point %d, %d\n", pointerToPoint.Name, pointerToPoint.X, pointerToPoint.Y)
}

func main() {

	// structs are value types
	houseOne := Point{"House One", 5, 2}
	houseTwo := houseOne
	houseTwo.Name = "House Two"
	houseTwo.X = 7
	houseTwo.Y = 8
	printPoint(houseOne)
	printPoint(houseTwo)

	// you can dereference pointers to structs to use
	// dot notation, but it's unnecessary (see print
	// function called below)
	houseThree := Point{"House Three", 3, 1}
	pointerToHouseThree := &houseThree
	printPointerToPoint(pointerToHouseThree)

	// you can initialize some values and other will
	// use default values
	townSquare := Point{Name: "Town Square"}
	printPoint(townSquare)

	// a struct's default value is the struct with the
	// default value of all its feilds
	var p Point
	printPoint(p)

}
