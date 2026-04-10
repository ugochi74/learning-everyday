package main

import (
	"fmt"
	"strings"
)

func main() {
	str := []string{"hello", "world", "(up)"}
	for i := 0; i < len(str); i++ {
		if str[i] == "(up)" {
			str[i-1] = strings.ToUpper(str[i-1])
			str = append(str[:i], str[i+1:]...)
			i--
		}
	}
	h := strings.Join(str, " ")
	fmt.Println(h)
}
