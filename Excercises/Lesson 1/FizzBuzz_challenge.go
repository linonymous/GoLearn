package Lesson_1
import "fmt"

func main()  {
	for x:=1; x<=20; x++{

		switch  {
		case x % 3 == 0 && x % 5 == 0:
			fmt.Println("FizzBuzz")
		case x % 3 == 0:
			fmt.Println("Fizz")
		case x % 5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(x)
		}

	}
}