/*
   Created by LINONYMOUS
   at 9/1/2019
*/
package Lesson_3

import (

	"fmt"
	"math"
)

type Rectangle struct{
	Length int
	Breadth int
}

func (s *Rectangle) Area() float64  {
	return float64(s.Length * s.Breadth)
}

type Circle struct {
	Radius float64
}
func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func sumAreas(shapes []Shape) float64  {
	total := 0.0

	for _, shape := range shapes {
		total += shape.Area()
	}

	return total
}

// an interface is a collection of methods
type Shape interface {
	Area() float64
}

func main()  {
	c := &Circle{Radius:3.5}
	r := &Rectangle{Length:9, Breadth:4}

	fmt.Println(r.Area())
	fmt.Println(c.Area())

	shapes := []Shape{r, c}

	sa:= sumAreas(shapes)
	fmt.Println(sa)
}