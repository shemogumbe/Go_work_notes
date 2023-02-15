package golearning

import "fmt"

func main() {
	book := "The color of magic"

	fmt.Println(book)

	fmt.Println(len(book))

	fmt.Println("book[0] = %v (Type %T)\n", book[0], book[0])

	end := book[4:]
	fmt.Println("Ommited first -", end)

	mid := book[4:10]
	fmt.Println("Omited first and last -", mid)

	start := book[:4]
	fmt.Println("Ommited last -", start)
}
