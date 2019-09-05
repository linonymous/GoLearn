/*
   Created by LINONYMOUS
   at 8/31/2019
*/
package Lesson_3

import "fmt"


type Trade struct {
	Symbol string
	Volume int
	Price float64
	Buy bool
}


//     receiver
func (t *Trade) Value() float64  {
	value := float64(t.Volume) * t.Price
	if t.Buy {
		value =- value
	}

	return value
}

func main()  {
	t := Trade{
		Symbol: "MSFT",
		Volume: 10,
		Price:  99.89,
		Buy:    true,
	}

	fmt.Println(t.Value())
}
