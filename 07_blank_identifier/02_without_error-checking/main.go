package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// _ means blank identifier and in this case means
	// don't check for an err because you would generally
	// put "err" in the place of the "_"
	res, _ := http.Get("http://www.mcleods.com")
	page, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s", page)
}
