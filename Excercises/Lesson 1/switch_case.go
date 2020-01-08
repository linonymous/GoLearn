package Lesson_1

import "fmt"

func main(){
	x := 1.0
	switch x {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("many")
	}

	switch  {

	case x > 100:
		fmt.Println("x is bigger than 100")
	case x > 10:
		fmt.Println("x is less than 10")
	default:
		fmt.Println("whattttt!")
	}
}
