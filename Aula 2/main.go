package main

import "fmt"

func main(){
	mensagem := "Xd"
	fmt.Println(mensagem)
	var a, b, c = true, "aLO", 21
	fmt.Println(a,b,c)

	var nome string = "Saullo Reis dos Santos"
	var existe bool = true

	if existe {
		fmt.Println(nome, existe)
	}
	fmt.Println(nome)

	var x, y = 10, 20
	fmt.Println(x, y)
	x, y = y,x
	fmt.Println(x,y)
}