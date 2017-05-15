package main

import "fmt"

//funcs are variadic, not arguments and params
//though this just shows how to pass variadic arguments
func main() {
	//slice of float64s
	data := []float64{43, 56, 87, 12, 45, 57}
	//the "..." here allows the float64 slice
	//to have each item within the slice
	//to be passed as an individual argument
	//instead of an entire slice
	//in order to give the function the parameter
	//type it's looking for
	n := average(data...)
	fmt.Println(n)
}

//this is not looking for a float64 slice, its looking
//for 0 through infinity arguments to be passed into it
func average(sf ...float64) float64 {
	total := 0.0
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}
