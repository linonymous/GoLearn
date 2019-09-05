/*
   Created by LINONYMOUS
   at 9/1/2019
*/
package Lesson_3

import (
	"fmt"
	"os"
)

type Point struct {
	X int
	Y int
}

func (P *Point) Move(dx int, dy int)  {
	P.X += dx
	P.Y += dy
}


type Square struct {
	Center Point
	Length int
}

func (s *Square) Area() int {
	return s.Length * s.Length
}

func NewSquare(x int, y int, length int) (*Square, error)  {
	if length <= 0 {
		return nil, fmt.Errorf("length must be greater than 0")
	}

	s := &Square{
		Center: Point{x, y},
		Length: length,
	}

	return s, nil
}

func (s *Square) Move(dx int, dy int)  {
	s.Center.Move(dx, dy)
}


func main()  {
	var square, err = NewSquare(3, 4, 19)

	if err != nil {
		os.Exit(1)
	}

	square.Center.Move(2, -1)
	fmt.Println(square)
	fmt.Println(square.Area())
}