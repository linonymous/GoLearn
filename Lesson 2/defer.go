/*
   Created by LINONYMOUS
   at 8/31/2019
*/
package Lesson_2

import (
	"fmt"
)

func cleanup(name string)  {
	fmt.Println("Cleaning up things! ", name)
}

func worker()  {
	defer cleanup("Swapnil")
	defer cleanup("omg")

	fmt.Println("worker")
}

func main()  {
	worker()
}
