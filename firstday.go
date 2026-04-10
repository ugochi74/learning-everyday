package main

import (
	"fmt"
)

func main() {
	for i := 'z'; i >= 'a'; i-- {

		fmt.Printf("%c", i)
		//fmt.Println()
	}
	fmt.Println()

	for i := 'a'; i <= 'z'; i++ {
		fmt.Printf("%c", i)
		//fmt.Println()
	}
	fmt.Println()

}
