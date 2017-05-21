package main

import "fmt"

//arrays are a number of elements
//of a specific type and cannot change their length(unlike slices)
//arrays are not commonly used
//they are an underlying datatype
//that are relied upon by other datatypes
//such as slices which are the more common option
func main() {
	//if you define a number
	//it's an array(limited amount of space)
	//if you leave the brackets empty
	//it's a slice
	var x [58]string
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[42])
	//65 is chosen because in ASCII
	//65 is the capital A character
	for i := 65; i <= 122; i++ {
		x[i-65] = string(i) //puts ascii num in position 0 through X index
	}
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[42])
}
