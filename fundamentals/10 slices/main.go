package main

import "fmt"

func main() {

	// slices are basically arrays but dynamically sized,
	// declare a slice literal like an array, but omit
	// the size value
	var primes []int = []int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// append to slice like so
	primes = append(primes, 17)
	fmt.Println(primes)

	// get the length of arrays/slices like so
	fmt.Println(len(primes))

	// get the capacity of arrays/slices like so
	fmt.Println(len(primes))

	// slices are really just references to arrays, changing a
	// slice modifies the underlying array
	names := [4]string{"John", "Ashley", "Tim", "Haley"}
	someNames := names[1:3]
	someNames[0] = "Joy"
	someNames = append(someNames, "Kate")
	fmt.Println(names)
	fmt.Println(someNames)

	// slice high low values are 0 and the length of the array,
	// so the following lines are the equivalent
	i := names[0:4]
	j := names[0:]
	k := names[:4]
	h := names[:]
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)
	fmt.Println(h)

	// slice default value is nil, it has a capacity of zero and
	// no underlying array
	var s []int
	fmt.Println(s)
	if s == nil {
		fmt.Println("nil")
	}

}
