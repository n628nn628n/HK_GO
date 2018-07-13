package main

import (
	"fmt"
	"strconv"

	"./otherpackage"
)

func main() {

	var a int = 1

	var b int32 = 2

	var c int64 = 3

	var d string = "999"

	var e float32 = 88.8

	var f float64 = 99.9

	var x string = "I Love Golang_"

	// a + b,用int型態印出
	fmt.Printf("a + b = (%T) %d \n", a+int(b), a+int(b))
	// a + b + c,用int型態印出
	fmt.Printf("a + b + c = (%T) %d \n", a+int(b)+int(c), a+int(b)+int(c))
	// f / e,用float型態印出
	fmt.Printf("f / e = (%T) %f \n", float32(f)/e, float32(f)/e)
	// a + d,用int型態印出
	i, _ := strconv.Atoi(d)
	fmt.Printf("f / e = (%T) %d \n", a+i, a+i)
	// x + a,字串相接,用string型態印出
	a1 := strconv.Itoa(a)
	a2 := fmt.Sprint(a)
	fmt.Printf("%s%s \n", x, a1)
	fmt.Printf("%s%s \n", x, a2)
	fmt.Printf("%s%d \n", x, a)

	// private func test
	fmt.Println(privateTest())
	otherpackage.OtherPackage()

}

// 除了main function之外,還要使用一個private function
func privateTest() string {
	return "private test"
}
