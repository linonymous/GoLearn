/*
   Created by LINONYMOUS
   at 9/1/2019
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

func NewTrade(symbol string, volume int, price float64, buy bool) (*Trade, error)  {
	if symbol == ""{
		return nil, fmt.Errorf("symbol can't be empty")
	}
	if volume <= 0 {
		return nil, fmt.Errorf("volume must be >= 0 (was %d)", volume)
	}
	if price <= 0.0 {
		return nil, fmt.Errorf("price must be >=0 (was %f)", price)
	}

	trade := &Trade{
		Symbol: symbol,
		Volume: volume,
		Price: price,
		Buy: buy,
	}

	return trade, nil
}

func main()  {

	my_trade, err := NewTrade("LINCP", 5, 12.34, true)
	if err == nil{
		fmt.Println(my_trade)
	}

}