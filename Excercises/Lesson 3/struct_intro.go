/*
   Created by LINONYMOUS
   at 8/31/2019
*/
package Lesson_3

import (
	"fmt"
)

type Trade struct {
	Symbol string
	Volume int
	Price float64
	Buy bool
}

// Everything that starts with upper case letter
// is accessible from other packages otherwise
// it is only accessible from same package

func main()  {

	trade1 := Trade{"MSFT", 10, 99.89, true}

	fmt.Println(trade1)

	fmt.Printf("%+v\n", trade1)

	fmt.Println(trade1.Symbol)

	t2 := Trade{
		Symbol: "MSFT",
		Volume: 10,
		Price:  99.89,
		Buy:    true,
	}

	fmt.Printf("%+v\n", t2)

	t3 := Trade{}
	fmt.Printf("%+v\n", t3)
}
