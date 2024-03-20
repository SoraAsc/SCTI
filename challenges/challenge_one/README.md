# Primeiro Desafio

Seu primeiro desafio em GO é simples, começe com a sua fork e faça checkout de um branch com `desafio_SEUNOME` então crie uma pasta aqui neste diretório com nome `SEUNOME_solucao`.

Dentro de sua pasta, inicialize um projeto em go com `go mod init pessoas` e crie o arquivo `desafio.go`, nele declare o pacote:

```go
package main

func main() {}
```

Seu desafio será a partir de um struct Pessoa que possui os campos `Nome`, `Idade` e `Cadastrado`, crie um slice `Pessoas` que ja possui duas pessoas:

`Nome: "teste", Idade: 33, Cadastrado: true`
`Nome: "Sophia", Idade: 22, Cadastrado: false`

Adicione mais 3 pessoas da sua escolha, o método não importa. Elas tem que ser adicionadas durante o código e não devem pré existir no slice (podem no código).

Então através do uso de um [Stringer](https://go.dev/tour/methods/17), escreva todas as pessoas no console usando só o objeto do struct como input do print, por ex:

Tenho uma pessoa chamada marcos:

`Marcos = {Nome: "Marcos", Idade: 12, Cadastrado: false}`

O meu print só podera ter `print(Marcos)`, o texto do stringer pode ser o que você quiser.

Se você tem alguma dúvida sobre como contribuir olhe [aqui](https://github.com/MintzyG/SCTI/blob/main/docs/CONTRIBUTING.md)

E qualquer dúvida em geral e direcionamento sinta-se livre para me mandar uma mensagem no privado :)

### Happy coding!
