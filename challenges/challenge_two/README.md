# Servidor HTTPS

O seu segundo desafio será uma expansão do primeiro.

Você irá expandir a sua aplicação para funcionar em um servidor HTTPS que serve JSON para o usuário.

O aplicativo terá a mesma base que o seu app atual ou seja:

- Um slice com dois usuários pre-existentes e um struct pessoa.

`Nome: Sophia, Idade: 21, Cadastrado: true, Id: 10`

`Nome: Teste, Idade: 32, Cadastrado: false, Id: 11`

Mas o app também terá funcionalidades diferentes, como:

- Usuários agora possuem Id
- Adicionar um novo usuário
- Listar todos os usuários
- Mostrar todos os usuários cadastrados

Todas essas respostas deverão ser retornadas em JSON.
Todas as requisições devem ser feitas através de CURL

O servidor terá os seguintes endpoints:

- `/app/user/create-user`

**O que fazer**

> Quando se receber a request para criar um novo usuário cria-lo em um id não existente maior que 10 e com os dados fornecidos pela request e adiciona-lo no slice

- `/app/user/list-users` Onde se retornam todos os usuários existentes em JSON

**O que fazer**

> Quando se receber essa request você deve retornar o slice inteiros de usuários em JSON

- `/app/user/registered-users`

**O que fazer**

> Quando se receber essa request você deve retorna uma lista com apenas os usuários cadastrados

## Materiais para ajuda:

- O que é [JSON](https://pt.wikipedia.org/wiki/JSON)
- O que é [CURL](https://pt.wikipedia.org/wiki/CURL)

### Curl no Linux:

```
curl -i(Inclui qual server) -X(Metodos) Tipo-Do-Metodo -H(Cabeçalho) "Dados do Cabeçalho" -d(Dados) "Dados da request em JSON" Server
```

**Exemplo de GET**

```
curl -X GET https://localhost:8080/list-users
```

**Exemplo de POST**

```
curl -i -X POST -H "Content-Type:application/json" -d '{"Nome": "Sophia",  "Idade": "21", "Cadastrado": true }" http://localhost:8080/create-user
```

Essa cURL request no endpoint `create-user` manda as informações do usuário a ser criado

### Curl no Windows

**Exemplo de GET**

```
Invoke-RestMethod -Uri "http://localhost:8080/list-users" -Method Get -Headers @{"Content-Type" = "application/json"}
```

**Exemplo de POST**

```
Invoke-RestMethod -Uri "http://localhost:8080/create-user" -Method Post -Headers @{"Content-Type" = "application/json"} -Body '{
    "Nome": "Sophia",
    "Idade": 21,
    "Cadastrado": true
}'
```

### Materiais para o exercício 

[Go http server](https://gobyexample.com/http-server)
[Retornar JSON](https://golangbyexample.com/json-response-body-http-go/)

Sinta-se livre para pedir ajuda para mim a qualquer momento :3


