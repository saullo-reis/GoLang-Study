package main

import (
	"fmt"
	"time"
)

func num(){
	for i := 0; i < 10; i++ {
		fmt.Printf("%d", i)
		time.Sleep(time.Millisecond * 150)
	}
}

func words() {
	for l := 'a'; l < 'j'; l++ {
		fmt.Printf("%c", l)
		time.Sleep(time.Millisecond * 230)
	}
}

func main() {
	go num()
	go words()
}