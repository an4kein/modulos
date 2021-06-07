package main

// this is example

import (
	"fmt"

	crtshgo "github.com/test"
)

func main() {
	test, err := crtshgo.GetJsonFromCrt("uber.com")
	if err != nil {
		fmt.Println("Error: ", err)
	}
	for _, v := range test {
		fmt.Println(v)
	}
	fmt.Printf("\ntype: >>>> %T\n", test)

}
