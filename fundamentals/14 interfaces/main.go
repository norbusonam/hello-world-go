package main

import "fmt"

type Any interface{}

// empty interfaces can be any type, all
// types have at least zero methods
func describe(thing Any) {
	fmt.Printf("%v has type %T\n", thing, thing)
}

// interfaces are sets of method signatures,
// no variables
type Abser interface {
	Abs() float64
}

type MyFloat float32

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Eater interface {
	Eat()
}

type Animal struct{}

func (a *Animal) Eat() {
	if a == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println("nom nom")
}

func main() {

	// interfaces are implemented implicitly,
	// we do not have to say that MyFloat
	// *implements* Abser, go figures that
	// out for you
	var f Abser = MyFloat(3.14159)
	fmt.Println(f)

	// empty interface demo
	describe(f)
	describe("cool string bro")

	// interfaces can themselved be nil
	var a Abser
	describe(a)

	// or their underlying values can be nil
	var eater Eater
	var animal *Animal
	eater = animal
	describe(eater)

	// interfaces with nil underlying values
	// can still be called
	eater.Eat()

	eater = &Animal{}
	eater.Eat()

}
