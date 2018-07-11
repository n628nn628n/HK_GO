package main

import (
	"fmt"
)

func main() {
	// 字串
	str := "ABC"
	// 字串轉[]Byte
	strToB := []byte(str)
	// []Byte轉字串
	bToStr := string(strToB)
	// 取字串中的一個字元
	oneOfStr := bToStr[0]
	// 硬出來
	fmt.Println(oneOfStr)

	// 數字
	num := 123
	// 轉64位元的數字
	num64 := int64(num)
	fmt.Printf("num 的型態 %T, 數值 %d or %v, 位址 %p\n", num, num, num, &num)

	// 數字轉字串
	numToStr := fmt.Sprint(num64)
	fmt.Printf("numToStr 的型態 %T, 數值 %d or %v, 位址 %p\n", numToStr, numToStr, numToStr, &numToStr)

	// 小數點
	floatNum := 12.34
	fmt.Printf("floatNum 的型態 %T, 數值 %d or %v, 位址 %p\n", floatNum, floatNum, floatNum, &floatNum)

	/**
	 * fmt.Printf 的 f 就是 format
	 * %d, %2d, %3d, ... 整數
	 * %f, %.2f, %.3f, ... 浮點數
	 * %s 字串
	 * %T 型態
	 * %p 記憶體位址
	 */

	// 數字陣列 (宣告變數時，即固定長度)
	numArray := [5]int{}
	fmt.Println("array", numArray)
	// 數字切片 (宣告變數時，無固定長度)
	numSlice := []int{}
	fmt.Println("slice", numSlice)
	// 數字切片 (宣告變數時，無固定長度，但一開始就會有5個元素)
	numSlice = make([]int, 5)
	fmt.Println("slice", numSlice)

	// Map
	aMap := map[string]int{}
	fmt.Println("map", aMap)
	aMap = make(map[string]int)
	fmt.Println("map", aMap)
}
