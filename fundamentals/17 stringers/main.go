package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("My name is %s and I am %d years old", p.Name, p.Age)
}

func main() {
	var me fmt.Stringer = Person{
		Name: "Norbu",
		Age:  22,
	}
	fmt.Println(me)
}
