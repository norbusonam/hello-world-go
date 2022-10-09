package main

import "fmt"

func main() {

	x := 10
	var y *int

	fmt.Printf("default pointer value:   %d (nil)\n", y)

	y = &x

	fmt.Printf("value of x:              %d\n", x)
	fmt.Printf("value of y:              %d\n", y)
	fmt.Printf("address of x (&x):       %d\n", &x)
	fmt.Printf("value at address y (*y): %d\n", *y)

}
