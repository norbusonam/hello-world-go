package main

import "fmt"

type Point struct {
	X, Y int
}

func main() {

	// defaults to nil, a map with no keys and
	// where no keys can be added
	var nilMap map[string]int
	fmt.Println(nilMap)
	if nilMap == nil {
		fmt.Println("nil")
	}

	// declare map literal and define key/value
	// pairs like so
	var priceMap = map[string]int{
		"iPhone 14 Pro": 999,
	}
	priceMap["iPhone 14"] = 799
	priceMap["iPhone 13"] = 699
	fmt.Println(priceMap)

	// get an element
	iPhone14Price := priceMap["iPhone 14"]
	fmt.Println(iPhone14Price)

	// delete a key/value pair
	delete(priceMap, "iPhone 13")
	fmt.Println(priceMap)

	// get an element with test
	iPhone13Price, ok := priceMap["iPhone 13"]
	if !ok {
		fmt.Println("\"iPhone 13\" is not a key in priceMap")
	} else {
		fmt.Println(iPhone13Price)
	}

	// if map types are structs, the type name can
	// be omitted
	var placeMap = map[Point]Point{
		{0, 0}: {1, 1},
		{1, 1}: {0, 0},
	}
	fmt.Println(placeMap)

	// maps are reference types
	p := placeMap
	delete(p, Point{0, 0})
	delete(p, Point{1, 1})
	fmt.Println(placeMap)
	fmt.Println(p)

}
