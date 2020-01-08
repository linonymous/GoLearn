/*
   Created by LINONYMOUS
   at 8/31/2019
*/
package Lesson_2

import (
	"fmt"
	"math"
)

func sqrt(n float64) (float64, error)  {
	if n < 0 {
		return 0.0, fmt.Errorf("SQRT of negarive value (%f)", n)
	}

	return math.Sqrt(n), nil

}

func main()  {
	s1, err := sqrt(2.0)
	if err!= nil{
		fmt.Println(err)
	}else {
		fmt.Println(s1)
	}

	s2, err := sqrt(-2.0)
	if err != nil{
		fmt.Println(err)
	}else {
		fmt.Println(s2)
	}
}