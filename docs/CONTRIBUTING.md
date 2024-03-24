# Como contribuir

Leia tudo antes de contribuir

Existem mais de uma maneira de contribuir para o projeto, uma é adicionando novas features e funcionalidades, outra é [criando Issues](#ISSUES) e corrigindo bugs, outra é escrevendo documentação, etc.
Escolha o que você se sente mais confortável fazendo e comece.

## Como começar:

1. [Faça uma fork](#FORK) desse repositório e clone para a sua máquina.
2. [Crie uma nova branch](#BRANCH) com o nome a apropriado.
3. Programe sua feature seguindo as [recomendações de código](#ISCLEAN) do repositório.
4. Quando tiver concluido [faça uma Pull-Request](#PR) para a branch `main` no repositório principar.

<a name="ISCLEAN"></a>
# Recomendação de padrão de código

Como não estamos trabalhando sozinhos é bom escrevermos códigos que todos os outros integrantes do projeto consigam entender facilmente, logo vamos tentar seguir certas convenções para facilitar o trabalho em grupo:

## Nomeando coisas
Evite nomes de variaveis e funções de poucas letras ou extremamente abreviados, tenha nomes de variáveis e funções descritiveis e que expliquem a natureza da variável ou função.

Entendemos que nomes que expliquem 100% da natureza de uma função por exemplo nem sempre são possíveis, mas não estamos pedindo por isso, só não nomeie sua funções e variáveis como o seguinte:

`fn proc(x, t, h) -> m`

Procure nomear as coisas para que a pessoa não fique 20 minutos lendo o código para entender o que ele faz.

## Comprimento das linhas

É possível que aconteça de alguma linha ficar extremamente grande, caso isso aconteça quebre a linha em múltiplas linhas, várias linguagens tem suporte para isso, e facilita a leitura do código imensamente, por exemplo:

#### Linha longa:

```Rust
  App::new().insert_resource(ClearColor(Color::rgb(0.1, 0.1, 0.1))).add_plugins((DefaultPlugins, PanCamPlugin::default())).add_systems(Startup, setup).run();
```

Difícil de ler e entender o que está acontecendo no código da linha

#### Multiplas linhas:

```Rust
  App::new()
    .insert_resource(ClearColor(Color::rgb(0.1, 0.1, 0.1)))
    .add_plugins((DefaultPlugins, PanCamPlugin::default()))
    .add_systems(Startup, setup)
    .run();
```

Mais fácil de rapidamente entender o que o código faz, ainda pode ser mais segmentada, mas isso já fica a critério do programador.

## Funções grandes

Funções ajudam muito a entender o código de maneira mais rápida mas isso deixa de ser verdade como uma função faz 50 coisas diferentes e tem 300 linhas, sabemos que algumas funções vão ser grandes por natureza, então não estamos pedindo para você só fazer funções de 10 linhas, mas só que evite que uma função faça coisas que poderiam ser feitas por mais de uma função e receba mais inputs do que o necessário.

Você também não precisa seguir o [Single Responsibility Principle (SRP)](https://en.wikipedia.org/wiki/Single_responsibility_principle), mas faça funções que não fiquem maiores e mais complicadas do que precisam ser.

## Evitando merge conflicts

Ninguém gosta quando um merge conflict acontece e trabalho pode ser perdido, então evite editar o mesmo arquivo ou trabalhar sem comunicação na mesma parte do código que algum outro programador, se for mexer na mesma feature que outra pessoa tanto na fork dela quanto em sua fork mas que modifica o mesmo arquivo, notifique e fique em contato com a pessoa que está trabalhando nesse arquivo para evitar merge conflicts.

## Faça commits com frequência

Comitar com frequência ajuda a manter um histórico de código simples e fácil de entender onde, caso aconteçam problemas pode se achar a causa facilmente e reverte-la, também ajuda a não perder progresso por motivos fora de seu controle.
Então sempre que fizer progresso no código, pare e faça um commit, por exemplo.

Você está programando uma função X, quando acabar faça um commit, fez a mudança Y faça um commit, mas tambem evite de fazer commits para coisas extremamente pequenas e poluir o histórico de commits.

Sempre dê nomes para seus commits que expliquem bem qual foi a mudança feita, lembre o nome do commit tem limite de 32 caracteres, caso você precise de mais texto para explicar a mudança use a descrição.

Exemplo:

Commit só com nome: `git commit -m "NAME"`

Commit com nome e descrição: `git commit -m "NAME" -m "DESCRIPTION"`

A descrição não tem limite de caracteres.

<a name="FORK"></a>
# Fazendo uma fork
- No canto superior direito da página deste repositório clique no botão `fork`
  
![image](https://github.com/MintzyG/SCTI/assets/21692264/f7b82130-3b79-4ac0-b5cb-0f04bbffe9c2)
- Mantenha todas as informações como estão e confirme a fork.
- A fork agora estará no seu perfil.
- Vá até o repositório no seu prefil e clone para o computador.

<a name="BRANCH"></a>
# Criando sua branch
- Depois de ter clonado o repositório use o seguinte comando no git:

`git checkout -b NAME`

## Nomeando sua branch
Como você deve nomear sua branch depende do que você irá fazer nela.

Se você vai programar uma feature nova o nome da branch deve ser _feat_ seguido do nome da feature separado por `/` por exemplo, se eu vou fazer a landing page o nome seria.

_feat:landing_page_ logo o comando para criar a branch seria `git checkout -b feat/landing_page`.

Por favor use snake_case na hora de dar nomes para suas branchs

#### Nomes a serem usados para branches:

- Se você estiver criando funcionalidade use `feat/NAME`
- Se você está corrigindo uma Issue/bug use `issue/NAME`
- Se você está escrevendo documentação use `doc/NAME`
- Se a mudança é puramente no design da aplicação use `design/NAME`

Sintam-se livres para sugerir mais.

## Como mudar de branch

Caso você tenha mais de uma branch e deseje mudar de branch use `git checkout NAME` sem o `-b`

Evite ter mais de uma feature branch para não se sobrecarregar.

<a name="ISSUES"></a>
# Criando e Contribuindo Issues

Achou um bug, erro, ou problema em algum lugar do código que não pode concertar agora? Tem alguma ideia ou sugestão de feature? Quer propor uma mudança?
Crie uma Issue com a tag apropriada para a situação e descreva bem o que quer, no caso de um bug/problema/erro descreva bem o erro, etapas para reproduzir, branch que você estava trabalhando quando o erro ocorreu,
possível causa do erro caso tenha alguma ideia.

<a name="PR"></a>
# Fazendo uma pull-request

Acabou a feature? concertou o bug? Faça o commit com todas as mudanças para a sua branch, e inicie uma pull request, nela explique tudo feito, removido e adicionado, mencione as Issues pertinentes com `#ISSUE_NUMBER` caso tenha mexido em alguma
e mande a PR para o repositório, antes dela ser aceita pelo menos 2 pessoa vão ter que fazer code-review e caso algum problema seja achado ela não será aceita até o mesmo ser concertado por quem autorou a PR. Após a PR não ter problemas, a mesma será colocada no repositório principal.
