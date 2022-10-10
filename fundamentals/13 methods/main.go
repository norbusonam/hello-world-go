package main

import "fmt"

// NOTE: Implementation of point is bad practice.
//       All methods on a type should be value or
//       pointer receivers, not a mixture of both!

// types must be declared in the same package in
// order to attach methods
type MyInt int

// non-structs can have methods
func (i *MyInt) Abs() {
	if *i < 0 {
		*i = *i * -1
	}
}

type Point struct {
	X, Y MyInt
}

// use pointer receivers to change the object,
// similarly to mutating *this* in other languages
func (point *Point) Abs() {
	point.X.Abs()
	point.Y.Abs()
}

// use non-pointer receivers to avoid mutating
// object
func (p1 Point) Add(p2 Point) Point {
	return Point{
		X: p1.X + p2.X,
		Y: p1.Y + p2.Y,
	}
}

func main() {

	p1 := Point{-5, 3}
	fmt.Println(p1)
	p1.Abs()
	fmt.Println(p1)
	p2 := Point{10, 10}
	p3 := p1.Add(p2)
	fmt.Printf("%v + %v = %v\n", p1, p2, p3)

	// go automatically implies resolved inderection
	p4 := Point{-12, -3}
	pointerToP4 := &p4
	p4.Abs()                  // (&p4).Abs()
	p5 := pointerToP4.Add(p4) // p5 := (*pointerToP4).Add(p4)
	fmt.Println(p4)
	fmt.Println(p5)

}
