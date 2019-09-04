/*
   Created by LINONYMOUS
   at 8/31/2019
*/
package Lesson_2

import "fmt"

func doubleAt(values []int, i int)  {
	values[i] *= 2
}

func doubleIt(value int)  {
	value *= 2
}

func doublePtr(value *int)  {
	*value *= 2
}

func main()  {
	values := []int{1, 2, 3, 4}
	doubleAt(values, 2)
	fmt.Println(values)
	doubleIt(values[3])
	fmt.Println(values)
	value := 4
	doubleIt(value)
	fmt.Println(value)
	doublePtr(&value)
	fmt.Println(value)
}