package main

import (
	"fmt"
	"strconv"
	"example.com/src/mypackage"
)

func Hello(name string) {
	fmt.Println("Hello", name)
}

func sum(a float32, b float32) float32{
	return a + b
}

func convertAndSum(a int, b string) (total int, err error){
	i, err := strconv.Atoi(b)
	if err != nil {
		return 
	}
	total = a + i
	return
}

func main(){
	Hello("Saullo")
	fmt.Println(sum(0.1, 0.2))
	pa.Teste()
	total, error := convertAndSum(1, "2")
	fmt.Println(total, error)
}