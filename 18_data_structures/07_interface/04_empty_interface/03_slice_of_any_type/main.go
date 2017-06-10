package main

import "fmt"

type animal struct {
	sound string
}

type dog struct {
	animal
	friendly bool
}
type cat struct {
	animal
	annoying bool
}

// this is bad practice because it breaks the type system
// which is needed to write consistent maintainable/predictable code
// Though things like fmt.Println take an empty interface type
// to be able to accept anything as a parameter - so there are some use cases
func main() {
	fido := dog{animal{"woof"}, true}
	garfield := cat{animal{"meow"}, true}
	critters := []interface{}{fido, garfield}
	fmt.Println(critters)
}
