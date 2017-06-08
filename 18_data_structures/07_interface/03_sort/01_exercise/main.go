package main

import (
	"fmt"
	"sort"
)

type people []string

func (p people) Len() int {
	return 2
}

func (p people) Swap(i, j int) {

}

func (p people) Less(i, j int) bool {
	return true
}

func main() {
	studyGroup := people{"Jay", "Silent-Bob", "Freddy", "Al"}
	sort.Sort(sort.StringSlice(studyGroup))
	// sort.Sort(sort.Reverse(studyGroup))
	fmt.Println(studyGroup)
}
