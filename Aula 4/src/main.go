package main

import (
	"fmt"
	"os"
	"log"
)

var (
	cara, coroa int
)

func lancarMoeda(lado string){
	switch lado {
	case "cara":
		cara++
	case "coroa" :
		coroa++
	default:
		fmt.Println("caiu em pé")
	}
}

func main(){
	a, b := 19, 13

	if a > b {
		fmt.Println("a maior que b")
	} else if a < b {
		fmt.Println("a menor que b")
	} else {
		fmt.Println("igual")
	}

	file , err := os.Open("hello.txt")
	if err != nil {
		log.Panic(err)
	}
	data := make([]byte, 100)
	if _, err := file.Read(data); err != nil {
		log.Panic(err)
	}

	fmt.Println(string(data))
}