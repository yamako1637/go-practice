package main

import (
	"fmt"
	"sync"
	"time"
)

func countDown(value string, limit int, wg *sync.WaitGroup) {
	fmt.Println(value + " Start")

	// 処理が終了したときに、並列処理が完了したことをwgに伝える
	defer wg.Done()

	for i := 0; i < limit; i++ {
		fmt.Println(value)
		time.Sleep(1 * time.Second)
	}
	fmt.Println(value + " End")
}

func main() {
	fmt.Println("main Start")
	var wg sync.WaitGroup

	// wg.Add(n)並列処理の個数を入れる
	wg.Add(4)

	// 並列処理
	go countDown("countDown1", 2, &wg)
	go countDown("countDown2", 3, &wg)
	go countDown("countDown3", 4, &wg)
	go countDown("countDown4", 5, &wg)

	// 上記の並列処理が終わるまで待機
	wg.Wait()
	fmt.Println("main End")
}
