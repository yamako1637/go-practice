package model

import (
	"fmt"
	"time"
)

type Index struct {
	Message     string `json:"message"`
	Count       int    `json:"count"`
	ProcessTime string `json:"process_time"`
}

func GetIndexMessage() Index {
	timeStart := time.Now()
	count := 0
	for i := 0; i <= 5000000; i++ {
		count = i
	}
	processTime := fmt.Sprintf("処理時間: %vms", time.Since(timeStart).Milliseconds())

	return Index{
		Message:     "Hello",
		Count:       count,
		ProcessTime: processTime,
	}
}
