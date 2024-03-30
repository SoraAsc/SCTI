package main

import "fmt"

type Pessoa struct {
	Nome       string
	Idade      int
	Cadastrado bool
}

func (p Pessoa) String() string {
	return fmt.Sprintf("%v, %v, %v", p.Nome, p.Idade, p.Cadastrado)
}

func main() {
	var s = []Pessoa{{"teste", 33, true}, {"Sophia", 22, false}}
	var Bruno Pessoa
	var Carol Pessoa
	var Emanuelito Pessoa

	fmt.Scanln(&Bruno.Nome, &Bruno.Idade, &Bruno.Cadastrado)
	fmt.Scanln(&Carol.Nome, &Carol.Idade, &Carol.Cadastrado)
	fmt.Scanln(&Emanuelito.Nome, &Emanuelito.Idade, &Emanuelito.Cadastrado)

	s = append(s, Bruno, Carol, Emanuelito)

	for _, p := range s {
		if p.Cadastrado {
			fmt.Println(p)
		}
	}
}
