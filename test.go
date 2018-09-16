package main

import (
	"fmt"
	"time"
)

func main() {

	timeStr := time.Now().Format("20060102")

	fmt.Println(timeStr)
}
