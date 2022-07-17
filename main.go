package main

import "time"

func main() {

	//x := 0

	//print("jss")

	for i := 0; i <= 100; i++ {

		go println(i)

	}

	time.Sleep(time.Second)
}
