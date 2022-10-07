package main

import "fmt"

func main() {

	// the following 3 lines perform the same assignment
	var x int = 1
	var y = 1
	z := 1
	fmt.Println(x, y, z)

	var i int    // default to 0
	var j string // default to ""
	var k bool   // default to false
	fmt.Println(i, j, k)

	const name = "Norbu" // constants can not be rename (obviously :P)
	fmt.Println(name)

}
