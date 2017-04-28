package main

import (
	"fmt"

	"github.com/AkyunaAkish/training/06_scope/01_package_scope/02_visibility/vis"
)

func main() {
	//vis.MyName isnt available unless the package variable
	//has a comment
	fmt.Println(vis.MyName)
	vis.PrintVar()
}
