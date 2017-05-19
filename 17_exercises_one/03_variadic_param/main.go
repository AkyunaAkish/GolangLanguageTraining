package main

import "fmt"

// write a func with one variadic param that finds
// the greatest number in a slice of numbers
func main() {
	a, b := findLargestParam()
	fmt.Printf("%T\n", a)                            //int
	fmt.Printf("%T\n", b)                            //bool
	fmt.Println(findLargestParam())                  //0 false
	fmt.Println(findLargestParam(-20, -10, 0, 4, 8)) //8 true
}

func findLargestParam(s ...int) (int, bool) {
	var largest int
	var ran bool

	if len(s) > 0 {
		ran = true
		largest = s[0] // set first item as largest
		s := s[1:]     // slice off first number

		for _, v := range s {
			if v > largest {
				largest = v
			}
		}
	} else {
		return 0, ran
	}

	return largest, ran
}
