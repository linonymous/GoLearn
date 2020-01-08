package Lesson_1

import "fmt"

func main()  {
	count := 0
	const MIN=1000
	const MAX=9999

	for a := MIN; a<= MAX; a++ {
		for b := a; b<= MAX; b++ {

			n := a * b
			s := fmt.Sprintf("%d", n)

			if s[0] == s[len(s) - 1] {
				count++
			}
		}
	}

	fmt.Println("Count is ", count)
}