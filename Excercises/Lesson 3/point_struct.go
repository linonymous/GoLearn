/*
   Created by LINONYMOUS
   at 9/1/2019
*/
package Lesson_3

import "fmt"

type Point struct {
	X int
	Y int
}
// pointer receiver
func (p *Point) Move(dx int, dy int)  {
	p.X += dx
	p.Y += dy
}

func main()  {
	p := Point{1, 2}
	p.Move(2, 3)
	fmt.Println(p)
}