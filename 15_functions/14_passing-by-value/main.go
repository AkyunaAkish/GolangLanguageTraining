package main

import "fmt"

//golang by default passes values into
//funcs by value with no reference
//to the original memory address
//if you want to affect the original
//value you can pass in the memory address
//and derefernce that value and reset it
//to change the original

//this is not true for reference data types
//slices, maps
//passing a slice or map to a function
//keeps reference to the memory address
//and if you change that value in one function
//it affects that data in another)
func main() {
	age := 44
	fmt.Println(&age)

	changeMe(&age)

	fmt.Println(&age)
	fmt.Println(age)
}

//int is passed in as memory address
//that stores an int
func changeMe(z *int) {
	fmt.Println(z)
	fmt.Println(*z)
	//*z dereferences memory address
	//and shows value instead of address
	*z = 24
	//changing value of original
	//data passed into func by using memory address
	fmt.Println(z)
	fmt.Println(*z)
}
