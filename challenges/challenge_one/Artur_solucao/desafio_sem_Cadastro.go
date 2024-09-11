package main

import (
	"fmt"
)

type Pessoa struct {
	Nome       string
	Idade      int
	Cadastrado bool
}

func (p Pessoa) String() string {
	return fmt.Sprintf("%v %v, %v", p.Nome, p.Idade, p.Cadastrado)
}

func main() {
	ab := []Pessoa{{"Sophia", 22, false}, {"teste", 33, true}}
	var Artur Pessoa
	var Bruno Pessoa
	var Zadoque Pessoa

	fmt.Println("Informe os dados: (Nome, Idade, Cadastrado)")
	fmt.Scanln(&Artur.Nome, &Artur.Idade, &Artur.Cadastrado)
	fmt.Println("Informe os dados: (Nome, Idade, Cadastrado)")
	fmt.Scanln(&Bruno.Nome, &Bruno.Idade, &Bruno.Cadastrado)
	fmt.Println("Informe os dados: (Nome, Idade, Cadastrado)")
	fmt.Scanln(&Zadoque.Nome, &Zadoque.Idade, &Zadoque.Cadastrado)
	fmt.Println("")

	ab = append(ab, Artur, Bruno, Zadoque)

	for _, Pessoa := range ab {
		if Pessoa.Cadastrado {
			fmt.Println(Pessoa)
		}

	}
}
