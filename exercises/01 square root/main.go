package main

import "fmt"

func abs(x float64) float64 {
	if x < 0 {
		return x * -1
	} else {
		return x
	}
}

func sqrt(x, threshold float64, maxIterations int) {
	// make guess
	z := 1.0
	fmt.Printf("Initial guess for sqrt(%f): %f\n", x, z)

	for i := 0; i < maxIterations; i++ {

		// compute
		zNew := z - (z*z-x)/(2*z)
		delta := abs(zNew - z)
		z = zNew

		// check threshold
		if delta < threshold {
			fmt.Printf("Threshold (%f) met at iteration %d\n", threshold, i+1)
			break
		}

		// logging for max iterations reached
		if i == maxIterations-1 {
			fmt.Printf("Maximum iteration (%d) reached\n", maxIterations)
		}
	}

	// log result
	fmt.Printf("Result for sqrt(%f): %f\n\n", x, z)
}

func main() {
	sqrt(100, 0.000001, 8)
	sqrt(69, 0.000001, 8)
	sqrt(2077, 0.000001, 8)
}
