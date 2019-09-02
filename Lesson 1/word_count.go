package Lesson_1

import (
	"fmt"
	"strings"
)

func main()  {

	text := `
		Needles and pins
		Needles and pins
		Sew me a sail
		To catch me the wind
	`

	count := make(map[string]int)
	words := strings.Fields(text)
	for _, word := range words{
		_, ok := count[word]
		if ok{
			count[word]++
		}else{
			count[word]=1
		}
	}

	fmt.Println(count)
}