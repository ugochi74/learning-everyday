package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "123)343) =78748="
	c := strings.Split(str, "")
	var g []string
	//fmt.Println(c)
	for i := 0; i < len(c); i++ {
		if c[i] == ")" || c[i] == " " {
			continue
		} else {
			g = append(g, string(str[i]))
		}
	}
	h := strings.Join(g, "")
	fmt.Println(h)
}
