package main

import (
	"fmt"
	"time"
)

func receice(name string, ch <-chan int) {
	for {
		i, ok := <-ch
		if ok == false {
			break
		}
		fmt.Println(name, i)
	}
	fmt.Println(name + " is done.")
}

func main() {
	ch := make(chan int, 20)

	go receice("routine 1", ch)
	go receice("routine 2", ch)
	go receice("routine 3", ch)

	i := 0
	for i < 10 {
		ch <- i
		i++
		time.Sleep(1 * time.Second)
	}
	close(ch)

	time.Sleep(1 * time.Second)
}
