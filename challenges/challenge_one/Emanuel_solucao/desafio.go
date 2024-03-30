package main

import "fmt"

type Pessoa struct {
    Nome string
    Idade int
    Cadastrado bool
}

func (p Pessoa) String() string {
  return fmt.Sprintf("Nome: %v, Idade: %v, Cadastrado: %v", p.Nome, p.Idade, p.Cadastrado)
}

func main() {
  Pessoas := []Pessoa{{Nome: "teste", Idade: 33, Cadastrado: true}, {Nome: "Sophia", Idade: 22, Cadastrado: false}}
  //fmt.Println("Informação contida no slice pessoas:", Pessoas)

  var NomeScan string
  var IdadeScan int
  var CadastradoScan bool

  for i := 1; i <= 3; i++ {
    fmt.Println("Nome da", i, "ª pessoa:")
    fmt.Scanln(&NomeScan)
    fmt.Println("Idade da", i, "ª pessoa:")
    fmt.Scanln(&IdadeScan)
    fmt.Println(i, "ª pessoa cadastrada? Escreva 1 para cadastrado e 0 para não cadastrado")
    fmt.Scanln(&CadastradoScan)

    Pessoas = append(Pessoas, Pessoa{NomeScan, IdadeScan, CadastradoScan})
}


  //fmt.Println("Informação contida no slice pessoas:", Pessoas)
  for _, Pessoa := range Pessoas{
    fmt.Println(Pessoa)
  }


}
