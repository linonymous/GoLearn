package Lesson_1

import "fmt"

func main()  {
	//var x float64
	//var y float64

	x, y := 1.0, 2.0
	fmt.Printf("x=%v, %T\n", x, x)
	fmt.Printf("y=%v, %T\n", y, y)

	x=1
	y=2

	fmt.Printf("x=%v, %T\n", x, x)
	fmt.Printf("y=%v, %T\n", y, y)

	var mean float64
	mean = (x + y) / 2.0
	fmt.Printf("Result= %v, %T\n", mean, mean)
}