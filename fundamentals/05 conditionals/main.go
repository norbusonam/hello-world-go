package main

import "fmt"

func main() {
	evenOrOdd(69)
	evenOrOdd(100)
	checkSumIsCorrect(5, 5, 10)
	checkSumIsCorrect(12, 32, 45)
}

// basic if else statement
func evenOrOdd(num int) {
	if num%2 == 0 {
		fmt.Printf("%d is even\n", num)
	} else {
		fmt.Printf("%d is odd\n", num)
	}
}

// if else statement with a short statement
func checkSumIsCorrect(num1, num2, userSum int) {
	if realSum := num1 + num2; userSum == realSum {
		fmt.Printf("That's correct, %d + %d = %d\n", num1, num2, realSum)
	} else {
		fmt.Printf("That's incorrect, %d + %d = %d, not %d\n", num1, num2, realSum, userSum)
	}
}
