package main
import (
	"fmt"
	"unsafe"
)

func main() {
	str := "odo"
	firstname := unsafe.StringData(str)
	fmt.Printf("Before: %s | Address %p\n", str, firstname)

	str  += " confidence"
	secondname := unsafe.StringData(str)
	fmt.Printf("After: %s | Address %p\n", str, secondname)
}