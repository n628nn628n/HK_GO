package main

import (
	"fmt"
	"strconv"
)

func main() {

	var a int = 1

	var b int32 = 2

	var c int64 = 3

	var d string = "999"

	var e float32 = 88.8

	var f float64 = 99.9

	var x string = "I Love Golang_"

	fmt.Printf("a + b = (%T) %d", a+int(b), a+int(b))
	fmt.Println("")
	fmt.Printf("a + b + c = (%T) %d", a+int(b)+int(c), a+int(b)+int(c))
	fmt.Println("")
	fmt.Printf("f / e = (%T) %f", float32(f)/e, float32(f)/e)
	fmt.Println("")
	i, _ := strconv.Atoi(d)
	fmt.Printf("f / e = (%T) %d", a+i, a+i)
	fmt.Println("")
	z := strconv.Itoa(a)
	fmt.Printf("%s%s", x, z)
	fmt.Println("")
	fmt.Printf("%s%d", x, a)
}
