package main

import "fmt"

func main() {
	//slice of slice of ints with a length starting with 0
	//but a capacity of 3
	transactions := make([][]int, 0, 3)

	fmt.Println(len(transactions)) //0
	fmt.Println(cap(transactions)) //3

	for i := 0; i < 4; i++ {
		transaction := make([]int, 0)
		for j := 0; j < 4; j++ {
			transaction = append(transaction, j)
		}
		transactions = append(transactions, transaction)
	}

	fmt.Println(transactions)
}
