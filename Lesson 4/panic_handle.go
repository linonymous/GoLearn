/*
   Created by LINONYMOUS
   at 9/1/2019
*/
package Lesson_4

import (
	"fmt"
)

func safeValue(vals []int, index int) int  {
	defer func() {
		if err := recover(); err != nil{
			fmt.Printf("ERROR: %s\n", err)
		}
	}()

	return vals[index]
}

// when an error cant be handled, you have to signal it
// panic function does that

func main()  {
	/*
	vals := []int{1, 2, 3, 4}

	fmt.Println(vals[10])*/
	/*
	file, err := os.Open("no-such-file")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	fmt.Println("File was opened.")
	*/

	v := safeValue([]int{1, 2, 3}, 10)

	fmt.Println(v)
}
