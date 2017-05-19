package main

import "fmt"

// write a function that takes an int
// the function will have 2 returns
// the first return should be the arg divided by 2
// the second return should be a bool telling if the num is even
func main() {
	fmt.Println(doubleReturn(5))
}

func doubleReturn(i int) (int, bool) {
	return i / 2, i%2 == 0
}

// with float64(if you want decimal values)
// func main() {
// 	fmt.Println(doubleReturn(5))
// }

// func doubleReturn(i int) (float64, bool) {
// 	return float64(i) / 2, i%2 == 0
// }
