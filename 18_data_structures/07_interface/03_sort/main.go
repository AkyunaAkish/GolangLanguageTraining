package main

import (
	"fmt"
	"sort"
)

// creating type with underlying []string type
// had to create Sort interface dependant methods
// so created Len Swap and Less from docs for Sort
type people []string

func (p people) Len() int {
	return len(p)
}

func (p people) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p people) Less(i, j int) bool {
	return p[i] < p[j]
}

func main() {
	studyGroup := people{"Jay", "Silent-Bob", "Freddy", "Al"}

	// the next two methods work even when giving Len,Swap, and Less
	// no real functionality(would work with regular type of []string rather than type people)

	// this version takes an interface which string slice is
	// sort.Sort(sort.StringSlice(studyGroup))

	// this method converts data to interface and calls
	// Sort method of that interface
	// sort.StringSlice(studyGroup).Sort()

	// reverse order
	// sort.Sort(sort.Reverse(studyGroup))
	// sort.Sort(sort.Reverse(sort.StringSlice(studyGroup)))
	// sort.Sort(sort.Reverse(sort.IntSlice(someIntSlice)))

	// this method works when giving Len,Swap, and Less
	// real functionality
	// if I just made a variable with a type of []string
	// rather than the type people
	// this method would also work
	sort.Sort(studyGroup)
	fmt.Println(studyGroup)

	// another way
	// s := []string{"a", "b", "d", "e"}
	// sort.Strings(s)

	// using the sort package
	// for StringSlice or IntSlice
	// you are converting your type
	// to ensure your data will implement
	// the correct interface methods
	// that is expeceted as an argument to sort

	// Reverse returns an interface
	// Sort needs an interface as an argument
	// which is why this is needed:
	// sort.Sort(sort.Reverse(sort.StringSlice(studyGroup)))

	// int examples
	// n := []int{1, 2, 3, 4, 7, 5, 10}
	// sort.Sort(sort.Reverse(sort.IntSlice(n)))
	// n := []int{1, 2, 3, 4, 7, 5, 10}
	// sort.Ints(n)
}
