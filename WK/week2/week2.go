package main

import (
	"fmt"
	"sort"
)

func main() {
	// 先建立一個長度為 5 的陣列來儲存這些測驗成績，接著將分數填入每個元素中。
	score := [5]int{60, 70, 80, 90, 100}

	// 再來使用一個迴圈來計算成績的總和。
	sum := 0
	for i := 0; i < 5; i++ {
		sum += score[i]
	}
	// 最後我們將成績的總和除以元素的數量，以取得平均值。
	fmt.Printf("Sum %v => Average %v \n", sum, sum/5)

	// 寫個程式找出串列中最小的數字：
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}
	sort.Ints(x)
	fmt.Printf("%d \n", x[0])

}
