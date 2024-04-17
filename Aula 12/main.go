package main

import (
	"fmt"
	"time"
)

func num(done chan <- bool){
	for i := 0; i < 10; i++ {
		fmt.Printf("%d", i)
		time.Sleep(time.Millisecond * 150)
	}
	done <- true
}

func words(done chan <- bool) {
	for l := 'a'; l < 'j'; l++ {
		fmt.Printf("%c", l)
		time.Sleep(time.Millisecond * 230)
	}
	done <- true
}

func main() {
	cn := make(chan bool)
	cw := make(chan bool)

	go num(cn)
	go words(cw)
	
	<-cn
	<-cw
}