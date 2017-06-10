package main

import "fmt"

// no methods defined so everything implements it
type vehicles interface{}

type vehicle struct {
	Seats    int
	MaxSpeed int
	Color    string
}

type car struct {
	vehicle
	Wheels int
	Doors  int
}

type plane struct {
	vehicle
	Jet bool
}

type boat struct {
	vehicle
	Length int
}

func main() {
	prius := car{}
	tacoma := car{}

	boeing747 := plane{}
	boeing757 := plane{}

	rides := []vehicles{prius, tacoma, boeing747, boeing757}

	for key, value := range rides {
		fmt.Println("key: ", key)
		fmt.Println("value: ", value)
	}

}
