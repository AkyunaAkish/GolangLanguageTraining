package main

import "fmt"

// A rune is:
// a character (single quotes instead of double or back tics)
// example: 'a'
// an integer value identifying a Unicode code point
// alias for int32
//  - how many bytes in 32 bits? (4 bytes --> 4 * 8 = 32)
//  - UTF-8 is a 4 byte coding scheme
//  - with int 32 (4 bytes) we have numbers for each of the code points
// UTF-8 is the most popular coding scheme
// which is used to encode characters
// for all languages and numbers

func main() {
	//conversion of string to slice of bytes
	//each character can be represented by 1-4 bytes
	//because UTF-8 is a 4 byte code scheme
	fmt.Println([]byte("Hello"))

	for i := 50; i <= 140; i++ {
		fmt.Println(i, " - ", string(i), " - ", []byte(string(i)))
	}
	foo := 'a'
	fmt.Println(foo)               //97 because that is the UTF-8/ascii number representation of the character a
	fmt.Printf("%T \n", foo)       //int32 is the type of the rune
	fmt.Printf("%T \n", rune(foo)) //int32 is the type of the rune converted to type rune explicitly

	// for i := 50; i <= 140; i++ {
	// 	fmt.Printf("%v - %v - %v \n", i, string(i), []byte(string(i)))
	// }
}
