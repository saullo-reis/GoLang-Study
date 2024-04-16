package main

import "fmt"

type Pessoa struct {
	Nome string
	Sobrenome string
	Idade uint8
}

func (p Pessoa) String() string {
	return fmt.Sprintf("Olá vai se %s e %d", p.Nome, p.Idade)
}

func (p Pessoa) name() string {
	return p.Nome
} 



func main(){
	 pessoa := Pessoa {
		Nome: "saullo",
		Sobrenome: "reis",
		Idade: 123,
	}

	fmt.Println(pessoa.String())
}