package main

import (
	"fmt"
	"strconv"
	"strings"
)

func applyUp(str []string, index int, count int) []string {
	cmd := ""
	num := 0
	for i := 1; i < len(str); i++ {
		if strings.HasPrefix(str[i], "(") {
			cmd = str[i]
			if i < len(str)-1 && strings.HasSuffix(str[i+1], ")") {
				temp, err := strconv.Atoi(str[i+1][:len(str[i+1])-1])
				if err != nil {
					return str
					//num = temp
				} //else {
				num = temp
			} else {
				num = 1
			}
			if num >= i {
				num = i
			}
			for j := i - num; j < i; j++ {
				switch cmd {
				case "(up,", "up)", "(up),":
					str[j] = strings.ToUpper(str[j])
				}
			}
		}
		if strings.HasPrefix(str[i], "(") {
			str = append(str[:i], str[i+1:]...)
			i--
		} else if strings.HasSuffix(str[i], ")") {
			str = append(str[:i], str[i+1:]...)
			i--
		}

	}
	return str

}

func main() {
	res := []string{"This", "is", "so", "exciting", "(up,", "2)"}
	fmt.Println(applyUp(res, 4, 3))
}
