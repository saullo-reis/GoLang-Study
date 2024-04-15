package main

import "fmt"

type Books struct {
	book_name string
	paginas uint8
}

type Pessoa struct {
	Nome string
	Sobrenome string
	Idade uint8
	Livros Books
}

func main(){
	//criação do map
	idades := make(map[string] uint8)

	//atribuicao no map
	idades["Tiago"] = 31
	idades["Joao"] = 32

	//Pegar se existe ou nao o valor no map
	val, ok := idades["Tiago"]

	if ok == false {
		panic("ERROR")
	}
	fmt.Println(val)
	books := Books {
		book_name: "Chimbinha",
		paginas: 80,
	}

	p := Pessoa {
		Nome: "Pessoa",
		Sobrenome: "Minde",
		Idade: 123,
		Livros: books,
	}

	fmt.Println(p)
}