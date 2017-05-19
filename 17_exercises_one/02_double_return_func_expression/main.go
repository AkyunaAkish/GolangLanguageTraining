package main

import "fmt"

// previous exercise as a func expression
func main() {
	doubleReturn := func(i int) (int, bool) {
		return i / 2, i%2 == 0
	}
	fmt.Println(doubleReturn(5))
}
