package main

import "fmt"

// structs

type Location struct {
	x, y, z int
}

type Human struct {
	name     string
	age      int
	location Location
}

type HumanError struct {
	what  string
	where Location
}

// stringers

func (l Location) String() string {
	return fmt.Sprintf("(%d, %d, %d)", l.x, l.y, l.z)
}

func (h Human) String() string {
	return fmt.Sprintf("%s is at %v", h.name, h.location)
}

// errors

func (err HumanError) Error() string {
	return fmt.Sprintf("Human Error: %s at %v", err.what, err.where)
}

// human functions

func (h *Human) fly(dz int) error {
	return HumanError{
		what:  "tried to fly",
		where: h.location,
	}
}

func (h *Human) walk(dx, dy int) error {
	if newX, newY := h.location.x+dx, h.location.y+dy; newX >= 0 && newY >= 0 {
		h.location.x = newX
		h.location.y = newY
		return nil
	}
	return HumanError{
		what:  "tried to go out of bounds",
		where: h.location,
	}
}

func main() {

	mark := &Human{
		name:     "Mark",
		location: Location{0, 0, 0},
	}
	fmt.Println(mark)

	// walk
	if err := mark.walk(2, 3); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(mark)
	}

	// walk out of bounds
	if err := mark.walk(-4, 0); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(mark)
	}

	// fly
	if err := mark.fly(10); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(mark)
	}

}
