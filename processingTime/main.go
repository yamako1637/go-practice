package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	count := 0
	// sum := 0
	timeStart := time.Now()
	// forで1億回処理
	for i := 0; i <= 1000000000; i++ {
		count = i
		// sum = sum + i
		if count%1000000000 == 0 && count != 0 {
			fmt.Printf(strconv.Itoa(count) + "\n")
		}
	}
	fmt.Printf("処理時間: %vms\n", time.Since(timeStart).Milliseconds())
	fmt.Printf("カウント: %v  \n", strconv.Itoa(count))
	// fmt.Printf("演算結果: %v  \n", strconv.Itoa(sum))
}
