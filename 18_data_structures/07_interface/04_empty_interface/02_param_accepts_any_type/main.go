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

// every type implements an empty interface
func specs(a interface{}) {
	fmt.Println(a)
}

func main() {
	fido := dog{animal{"woof"}, true}
	garfield := cat{animal{"meow"}, true}
	specs(fido)
	specs(garfield)
}
