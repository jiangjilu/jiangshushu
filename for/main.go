package main

import (
	"fmt"
	"time"
)

func main() {
	a := 0
	for {
		a++
		fmt.Println(a)
		time.Sleep(time.Second * 1)
	}
}
