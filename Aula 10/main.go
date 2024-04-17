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
	Pessoa
	RazaoSocial string
	cnpj string
}

func (pf PessoaFisica) Doc() string {
	return fmt.Sprintf("meu cpf é %d", pf.cpf)
}

func (pj PessoaJuridica) Doc() string {
	return fmt.Sprintf("meu cpf é %d", pj.cnpj)
}

type Document interface {
	Doc() string
}

func show(doc Document) {
	switch doc.(type) {
	case PessoaFisica :
		fmt.Println(doc.(PessoaFisica).Sobrenome)
	case PessoaJuridica:
		fmt.Println(doc.(PessoaJuridica).RazaoSocial)
	default:
		fmt.Println("Tipo desconhecido")
	}
	
	fmt.Println(doc)
	fmt.Println(doc.Doc())
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


	show(pessoa)

	pessoaJurid := PessoaJuridica { 
		Pessoa{
		Nome: "saullo",
		Idade: 123,
		Status: true,
		}, 
		"Reis", 
		"001.000.000-30",
	}

	show(pessoaJurid)
}