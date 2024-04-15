package main

import "fmt"

func main(){
	nomes := []string{"Tiago", "Daniel", "Joao"}
	var i int

	for i < len(nomes){
		fmt.Println(nomes)
		i++
	}
	for k, name := range nomes{
		fmt.Println(k, name)
	}
}