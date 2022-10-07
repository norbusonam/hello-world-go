package main

import "fmt"

func main() {

	forLoop(2)
	whileLoopEquivalent(2)
	// infiniteLoop()

}

func forLoop(count int) {
	for i := 0; i < count; i++ {
		fmt.Println(i)
	}
}

func whileLoopEquivalent(count int) {
	i := 0
	for i < count {
		fmt.Println(i)
		i++
	}
}

func infiniteLoop() {
	for {
		fmt.Println("don't call me")
	}
}
