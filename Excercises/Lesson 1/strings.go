package Lesson_1



import (
	"fmt"
)

func main()  {
	book := "The colour of magic"

	fmt.Println(book)
	fmt.Println(len(book))

	fmt.Printf("Book[0] := %v (type %T)\n", book[0], book[0])
	// uint8 is a byte
	// strings in go are immutable
	// book[0]=116 is not a valid assignment

	//strings can be sliced


	fmt.Println(book[4:11])
	fmt.Println(book[4:])
	fmt.Println(book[:4])


	// use + sign to concatenate strings
	fmt.Println("t" + book[1:])

	// Unicode
	fmt.Println("It was 1/2 price!")

	// Multi line strings
	poem := `
		The road goes ever on,
		down from the door where it began
		...
	`

	fmt.Println(poem)

}