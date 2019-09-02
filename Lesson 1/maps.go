package Lesson_1

import "fmt"

func main()  {

	stocks := map[string]float64{
		"AMZN": 1198.2,
		"GOOG": 1129.19,
		"MSFT": 98.61,
	}

	fmt.Println(len(stocks))
	fmt.Println(stocks["GOOG"])
	stocks["AMZN"] = 543.1
	fmt.Println(stocks["AMZN"])

	// safe access of the value in maps
	value, ok := stocks["MSFT"]
	if !ok {
		fmt.Println("TSLA not found")
	} else {
		fmt.Println(value)
	}

	for key, value := range stocks {
		fmt.Println(key, " : ", value)
	}
}
