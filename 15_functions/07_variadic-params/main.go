package main

import "fmt"

//the final parameter in a function signature
//may have a type prefixed with "..."
//a function with such a parameter is called
//"variadic"
//and may be invoked with zero or more arguments
//for that parameter

func main() {
	n := average(43, 56, 87, 12, 45, 57)
	fmt.Println(n)
}

func average(sf ...float64) float64 {
	total := 0.0
	// var total float64 //equivalent to line 18 because
	//0.0 is the 0 value for type float64
	fmt.Println(sf)
	fmt.Printf("%T \n", sf)
	//range allows you to loop over
	//slices(arrays)
	//ind is the index, and v is the value
	//if you don't need the index use "_" as a blank identifier
	for ind, v := range sf {
		fmt.Println("ind: ", ind, "value: ", v)
		total += v
	}
	//returns the average
	return total / float64(len(sf))
}
