package main

import "fmt"

type Pessoa struct {
	Nome string
	Idade uint8
	Status bool
}

type PessoaFisica struct {
	Pessoa
	Sobrenome string
	cpf string
}

type PessoaJuridica struct {
	RazaoSocial string
	cnpj string
}

func main(){
	pessoa := PessoaFisica { 
		Pessoa{
		Nome: "saullo",
		Idade: 123,
		Status: true,
		}, 
		"Reis", 
		"000.000.000-30",
	}


	fmt.Println(pessoa)
}