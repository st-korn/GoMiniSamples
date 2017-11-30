package main

import (
	"fmt"
	"time"
)

func counter() {
	for i:=0; i<5; i++ {
		fmt.Println(i)
		time.Sleep(time.Millisecond*100)
	}
}

func main() {
	go counter()
	time.Sleep(time.Millisecond*200)
	fmt.Println("Hello, world!")
	time.Sleep(time.Millisecond*500)
}