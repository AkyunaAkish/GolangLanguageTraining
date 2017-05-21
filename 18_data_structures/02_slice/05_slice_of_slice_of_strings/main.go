package main

import "fmt"

func main() {
	//slice which stores a slice of string
	//to add more dimensions to a multidimensional slice:
	//just add more: [][][][][]string
	records := make([][]string, 0)
	//student 1
	student1 := make([]string, 4)
	student1[0] = "Forest"
	student1[1] = "Griffin"
	student1[2] = "100.00"
	student1[3] = "74.00"
	//store the record
	records = append(records, student1)
	//student2
	student2 := make([]string, 4)
	student2[0] = "Amanda"
	student2[1] = "Bynes"
	student2[2] = "92.00"
	student2[3] = "96.00"
	//store the record
	records = append(records, student2)
	//print
	fmt.Println(records) // [[Forest Griffin 100.00 74.00] [Amanda Bynes 92.00 96.00]]
}
