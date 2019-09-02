package Lesson_1

import "fmt"

func main()  {
	nums := []int{1, 2, 3, 4, 5, 6}

	max := nums[0]
	for _, value := range nums[1:]{
		if value > max {
			max = value
		}
	}

	fmt.Println(max)
}
