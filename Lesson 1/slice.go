package Lesson_1

import (
	"fmt"
)

func main()  {
	// slices
	loons := []string{"bugs", "daffy", "taz"}
	fmt.Printf("loons = %v (type %T)", loons, loons)


	fmt.Println(len(loons))

	fmt.Println(loons[1])

	fmt.Println(loons[1:])

	var a []int

	fmt.Println(a)

	for i := 0; i< len(loons) ; i++ {
		fmt.Println(loons[i])
	}

	for i := range loons{
		fmt.Println(i)
	}

	for i, name := range loons{
		fmt.Println(i, name)
	}

	for _, name := range loons {
		fmt.Println(name)
	}

	loons = append(loons, "swapnil")
	fmt.Println(loons)


	//var slice []int= make([]int, n)


}
