package main

import "fmt"

// basic if else statement
func evenOrOddWithIfElse(num int) {
	if num%2 == 0 {
		fmt.Printf("%d is even\n", num)
	} else {
		fmt.Printf("%d is odd\n", num)
	}
}

// basic switch case statement
func evenOrOddWithSwitchCase(num int) {
	switch num % 2 {
	case 0:
		fmt.Printf("%d is even\n", num)
	case 1:
		fmt.Printf("%d is odd\n", num)
	default:
		fmt.Println("math is broken, this should never happen!")
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

func main() {
	evenOrOddWithIfElse(69)
	evenOrOddWithSwitchCase(100)
	checkSumIsCorrect(5, 5, 10)
	checkSumIsCorrect(12, 32, 45)
}
