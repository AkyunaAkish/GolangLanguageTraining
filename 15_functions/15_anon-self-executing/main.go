package main

import "fmt"

func main() {
	func() {
		fmt.Println("Im executed immediately")
	}()
}
