package Lesson_2

import "fmt"

func add(a int, b int) int  {
	return a + b
}

func divmod(a int, b int) (int, int)  {
	return a / b, a % b
}

func main()  {
	val := add(1, 2)
	fmt.Println(val)

	div, mod := divmod(7, 2)
	fmt.Println(div, ":", mod)

	div_a,_ := divmod(8, 4)
	fmt.Println(div_a)
}