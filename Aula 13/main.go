package main

import (
	"fmt"
	"time"
)

func num(n chan <- int){ //declara o channel
	for i := 0; i < 10; i++ {
		n <- i
		fmt.Printf("Escrito no channel %d\n", i)
		time.Sleep(time.Millisecond * 150)
	}
	close(n) //fecha o channel
}


func main() {
	cn := make(chan int) // criação do channel
	//cexemplo :-= make(chan int, 3) isso aqui abre 3 espaços para esse channel isso é o buffer
	go num(cn)
	
	for v := range cn{ //leitura do channel
		fmt.Printf("Lido do channel: %d\n", v)
	}
}