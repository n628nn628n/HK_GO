package main

import (
	"fmt"
	"log"
)

func main() {
	// wg := new(sync.WaitGroup)
	defer fmt.Println("Over1")

	defer func() {
		catchErr := recover()
		if catchErr != nil {
			log.Println("catch ", catchErr)
		}
	}()

	// ch := make(chan int)
	i := 0
	a := 1 / i // ---> panic
	fmt.Println("Hi", a)
	fmt.Println("Hi")

	// db := Connection()
	// defer db.Close()

	defer fmt.Println("Over2")

	fmt.Println("Hi")
	fmt.Println("Hi")
	// for i := 0; i < 1; i++ {
	// 	// wg.Add(1)
	// 	log.Println("Start")
	// 	go func(i int) {
	// 		// time.Sleep(time.Second * 3)
	// 		ch <- i
	// 		// fmt.Println(i)
	// 		// wg.Done()
	// 	}(i)
	// }

	// wg.Wait()
	// go print()
	// runtime.Gosched()
	// fmt.Println("AAA")
}

func print() {
	defer fmt.Println("Over3")
	fmt.Println("Hello")
}
