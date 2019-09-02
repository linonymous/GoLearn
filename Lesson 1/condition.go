package Lesson_1

import "fmt"

func main()  {
	x := 10

	if x > 5 {
		fmt.Println("X is big")
	}

	if x > 100 {
		fmt.Println("X is very big")
	}else {
		fmt.Println("X is not very big")
	}

	if x > 5 && x < 15 {
		fmt.Println("X is just right!")
	}

	if x < 20 || x > 30 {
		fmt.Println("X is out of range!")
	}

	a := 11.0
	b := 20.0

	if frac := a / b; frac > 0.5 {
		fmt.Println("a is more than half of b")
	}
}
