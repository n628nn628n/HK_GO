package main

import (
	"fmt"
)

func main() {
	a := "string"
	b := 123
	c := true
	d := []int{
		1,
		2,
		3,
	}
	e := map[int]string{
		1: "aa",
		2: "bb",
	}
	fmt.Println(a, b, c, d, e)

	var f interface{}
	f = 1
	fmt.Println("f", f)
	f = "abc"
	fmt.Println("f", f)

	type Person struct {
		name string
		age  int
	}

	me := Person{
		name: "sera",
	}

	fmt.Println(me.age)

	people := []Person{
		me,

		{"day", 100},

		{
			name: "lee",
			age:  180,
		},
	}
	var people2 []struct {
		age int
	}

	fmt.Println(people, people2)

	hello()

}

func hello() {
	fmt.Println("Hello")
}
