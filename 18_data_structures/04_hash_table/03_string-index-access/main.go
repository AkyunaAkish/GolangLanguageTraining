package main

import "fmt"

func main() {
	word := "Hello"
	fmt.Println(string(word[0])) // H
	letter := rune(word[0])
	fmt.Println(letter) // 72
}
