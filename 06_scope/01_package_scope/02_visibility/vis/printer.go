package vis

import "fmt"

//PrintVar exported function
func PrintVar() {
	//available in package vis and outside
	fmt.Println(MyName)
	//only available within package vis
	fmt.Println(yourName)
}
