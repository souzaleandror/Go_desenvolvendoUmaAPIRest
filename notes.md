#### 29/09/2023

Curso de Go: desenvolvendo uma API Rest

```
export GO111MODULE="on" 

Error = package github.com/lib/pq is not a main package
go env -w GO111MODULE=off
go get github.com/lib/pq
go mod init

lsof -i:8080
Kill -9
```

@01-Json, Request, Response e Go

@@01
Apresentação

[00:00] Olá, meu nome é Guilherme Lima e eu queria te dar as boas-vindas neste treinamento de API REST utilizando a linguagem GO. Neste curso vamos criar uma API REST utilizando o GO para guardar personalidades.
[00:15] Como assim personalidades? Na minha rua, na minha cidade, existem vários nomes de pessoas que eu não faço a mínima ideia de quem são e eu quero guardar essas informações. Eu coloquei dois nomes aqui que vocês estão vendo na tela de nomes de ruas da minha cidade. Eu coloco o nome dessa personalidade e um pouco da história dela.

Deodato Petit Wertheimer: Deodato Petit Wertheimer foi um médico e político brasileiro, seus primeiros anos de vida foram em São Paulo, mas logo mudou para Nova Friburgo no Estado do Rio de Janeiro e com 11 anos de idade ingressou no Colégio Anchieta dos jesuítas.
Carmela Dutra: Carmela Teles Leite Dutra foi a primeira-dama do Brasil, de 31 de janeiro de 1946 até a sua morte, tendo sido a esposa de Eurico Gaspar Dutra, 16º Presidente do Brasil. Era, carinhosamente, chamada de Dona Santinha, pela sua forte religiosidade, fazendo seu marido abrir uma capelinha no Palácio Guanabara.
[00:37] Para criar essa aplicação vamos criar um API REST do zero utilizando o GO. Vamos desde os princípios, como retornamos uma requisição com JSON, como fazemos o roteamento desses endpoints, como conectamos com banco de dados.

[00:55] Um ponto legal neste curso é que vamos usar o banco de dados no Docker, vamos conectar a nossa API REST consumindo um banco de dados no Docker. Vamos usar o RM do GO para conseguir conectar com o banco de dados e manipular mais fácil as nossas requisições sem precisar ficar escrevendo SQL a mão.

[01:15] Vamos criar o CRUD completo, vamos ser capazes de criar recursos, deletar, editar, atualizar. Serão muitas coisas legais. Além disso, vamos pensar em como deixamos a nossa API com um nível bom de estrutura, informando qual vai ser o Content-Type que estamos devolvendo, vamos criar middlewares também e vamos integrar a nossa API com o React.

[01:40] Ao final do curso vamos entender, depois que criamos todo nosso esqueleto, a nossa API como integramos o nosso back-end GO com um front-end React, que é o que vocês estão vendo aqui. Tanto que temos o símbolo do GO e o símbolo do React girando aqui em cima.

[01:56] Esse curso é indicado para as pessoas que já fizeram os treinamentos de GO da Alura e vão dar continuidade. Se você nunca fez nenhum treinamento de GO eu aconselho que você faça primeiro os cursos de GO descritos nessa formação e depois você siga para esse curso de API REST.

[02:14] Se você já tem familiaridade com outra linguagem de programação e você quer desenvolver uma API utilizando o GO, você é bem-vindo nesse curso também. Esse curso não é indicado para as pessoas que já tem uma maturidade muito grande na linguagem GO, se você já trabalha com GO, você já tem uma maturidade muito avançada, você já domina esse tema esse conteúdo, talvez esse curso não seja indicado para você.

[02:36] Sabendo disso vamos começar a nossa API de personalidades.

@@02
Preparando o ambiente

Olá!
É muito bom receber você neste curso de Criando uma API Rest com Go.

Espero que seja uma experiência de aprendizado incrível e que possamos vencer todos os desafios juntos. Neste curso, vamos aprofundar nosso conhecimento utilizando Go e criar uma API Rest do zero, seguindo boas práticas de programação. Além disso, vamos aprender também como usar um ORM no Go.

Preparando ambiente
Para conseguir acompanhar este curso, é recomendado que você tenha o Go instalado.
Durante o treinamento, usarei o VSCode para editar o código e para tornar o desenvolvimento ainda mais fácil, recomendo a instalação da extensão do Go desenvolvida pelo time da Google.

A Alura é formada por pessoas que gostam de tecnologia e acreditam no poder da educação através dela. Somos uma comunidade que ama compartilhar conhecimento. Em caso de dúvida na instalação ou durante o curso, conte sempre com o fórum. Caso não tenha dúvidas, não deixe de participar do fórum para ajudar outras pessoas e fazer da comunidade um lugar ainda melhor!
:)

https://www.youtube.com/watch?v=x39vqeBTUmE

https://golang.org/

https://code.visualstudio.com/

https://marketplace.visualstudio.com/items?itemName=golang.go

@@03
Iniciando o projeto

[00:00] Vamos iniciar os nossos estudos. Antes, para começarmos um projeto GO colocávamos esse projeto na GOPATH, eu vou abri-la para vocês verem. Com base nos nossos cursos anteriores eu vou abrir aqui, tínhamos no meu usuário uma pasta chamada “GO” que tinham três arquivos e criamos esses arquivos dentro de GitHub, começamos nosso projeto dentro desse GOPATH.
[00:25] Hoje existe uma forma de conseguirmos iniciar o nosso projeto em qualquer pasta e eu vou mostrar isso para vocês. Como vamos iniciar o projeto do zero eu vou clicar com o botão direito do mouse e vou clicar em "Nova pasta". Eu vou chamar essa pasta de “api-go-rest”, essa vai ser a pasta do nosso desenvolvimento.

[00:47] Vou usar o Visual Studio Code, como editor de código e vou arrastar essa “api-rest” aqui para o VS Code. Ele vai abrir e agora eu vou usar o terminal integrado do VS Code para executarmos o código que vai nos permitir que rodamos qualquer projeto GO sem que estejamos alocados na GOPATH.

[01:13] Vou utilizar as teclas "Ctrl + J" ou “Command + J" se você estiver utilizando MAC. O comando será assim: go mod init e agora vou colocar o comando init e vou dar um nome para esse projeto que vamos criar. Geralmente vinculamos esse projeto a algum repositório que temos lá no GitHub para hospedarmos o nosso código, essa é uma ideia.

[01:39] Você poderia colocar aqui link do GitHub, por exemplo: github.com/guilhermeonrails que é o meu usuário do GitHub, e eu coloco o nome desse projeto, voo colocar: go-rest-api, por exemplo. api-go-rest >>> go mod inint github.com/guilhermeonrails/go-rest-api.

[01:53] Se você não tem um repositório que você queira vincular com o código que vamos desenvolver aqui, se você não tem o GitHub, não tem problema você pode colocar o nome que você quiser. Por exemplo: só go-rest-api.

[02:05] Eu vou deixar esse nome mesmo, vou dar um "Enter" e ele fala que foi criado do go.mod. Vamos ver o que isso significa, quando abrimos a nossa pasta temos esse arquivo “go.mod” e ele fala qual é o módulo que estamos utilizando e onde todos os nossos pacotes e bibliotecas relacionadas ao GO estarão e ele também fala qual a versão GO estamos utilizando.

[02:29] Dessa forma conseguimos iniciar o nosso projeto GO sem que ele seja vinculado com a GOPATH. Para testarmos isso de fato eu vou criar um arquivo chamado “main.go” e nesse arquivo eu vou falar qual é o pacote que ele pertence package main, coisas que nós já vimos nos cursos anteriores de Orientação a Objetos, de GO com Web.

[02:52] Falei qual é o pacote que ele pertence. Vou criar uma função que vou chamar de func main() {. Observe que estou digitando e as coisas já estão acontecendo, já estão aparecendo aqui. Isso acontece porque estou usando uma extensão que é mantida pelo próprio GOogle.

[03:08] Se eu digitar aqui só go e der um "Ente", nesse bloco aqui de extensões ele fala que esse primeiro módulo estou usando, tanto que ele está até habilitado. Com ele conseguimos ter vários atalhos, tudo isso que eu digitei aqui, o package main e a func main(), eu poderia ter feito assim: pkgm e quando eu dou um "Enter" ele coloca qual é o pacote main e coloca a função também.

[03:31] Isso vai auxiliar no nosso desenvolvimento, vamos conseguir podar de forma mais rápida também. Se você não tem essa extensão eu recomendo que você instale para o VS Code.

[03:42] Nesse treinamento eu faço bastante menção do VS Code, eu sugiro até que você utilize esse mesmo editor que eu. Se você usa o Atom ou outro editor, você pode procurar também essas extensões do GO para o seu editor que você estiver usando.

[03:58] Criei aqui o “main.go” que por enquanto não faz nada e vou exibir uma mensagem simples: fnt.Println e repare que quando eu dei o "Enter" ele já trouxe aqui e fez o import para mim. Se não apareceu o import você pode digitá-lo import “fnt”. Vou colocar: ("Iniciando o servidor Rest com GO"). Claro que essa é uma mensagem fictícia não estamos iniciando neste caso o servidor com GO, vou rodar esse nosso projeto com o go run main.go e quando eu dou "Enter" ele fala: "Iniciando o servidor Rest com GO". Exibiu a mensagem e deu tudo certo.

package main

import "fmt"

func main() {
    fmt.Println("Iniciando o servidor Rest com Go")
}COPIAR CÓDIGO
[04:40] Só que só essa mensagem é um pouco simples, vamos abrir uma página usando outros pacotes e bibliotecas que precisamos. Vou criar uma função que vou chamar de Home() que vai representar a nossa página inicial. func Home(w http.ResponseWritter, r *http.Request). E como já vimos no curso de Web, sempre quando vamos criar uma nova função com o GO relacionada a abrir uma página, precisamos receber a requisição e responder utilizamos dois atalhos que são esses: ResponseWriter e o Request. O w para o responsewriter e o r apontando para o request que vamos utilizar.

[05:24] E vou colocar uma mensagem só não vou fazer nada mais que isso. fnt.Fprint(w, “Home Page”) porque vamos executar uma função que vamos passar o nosso ResponseWriter, responsável por escrever na tela, e aqui vou passar uma mensagem vou colocar, por exemplo, "Home Page".

[05:41] Quando eu salvo isso repare que ele trouxe para mim esse outro módulo "net/http", que será responsável por receber as requisições. É um módulo que vai nos auxiliar nessa parte de quando receber um requisição nesse endereço faça tal coisa e mensagem, ele vai auxiliar nessa conversa de cliente servidor também.

package main

import (
    "fmt"
    "net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Home Page")
}

func main() {
    fmt.Println("Iniciando o servidor Rest com Go")
}COPIAR CÓDIGO
[06:06] Vamos executar isso daqui, vamos colocar essa função Home() para funcionar. Para isso eu vou criar outra função que vai ser responsável por todas as rotas que temos na nossa aplicação. Quando navegamos e entramos em um site esse site tem várias páginas, cada página tem sua respectiva função que vai falar: "quando carregar essa página você vai abrir isso, isso e isso".

[06:31] Essa função temos o hábito de chamar no GO de func HandleResquest(). “Handle” em inglês significa “lidar”, HandleResquest é lidar com as requisições, algo desse tipo. Eu vou colocar aqui usando o módulo que já temos que é o: http.HandleFunc(), eu vou falar assim: "existe uma função que tem esse endereço e vai ser atendida por alguém."

[06:58] Sempre que chegar uma requisição no barra "/", ou seja, quando colocarmos o endereço da nossa página sem nada, sem "/compras" ou "/produtos" eu quero que ele caia nessa página específica e quem vai ser responsável por atendê-lo será a nossa função Home(), só isso, http.HandleFunc(“/”, Home).

[07:15] Agora aqui no nosso arquivo “main” eu vou executar esse nosso HandleResquest() executando ele aqui. Vamos relembrar o que fizemos: criamos uma função chamada Home() e falamos que toda vez que chegar uma requisição em Home() você vai responder exibindo essa mensagem "Home Page". Depois criamos alguém para lidar com essas requisições, quando chegar uma requisição em barra quem vai atender será o controle Home.

[07:41] Na nossa função main() falamos assim: "HandleResquest, executa", ele vai ficar "ouvindo" isso daqui. Só que se ele vai estar ouvindo o nosso servidor precisa estar funcionando, vou subir o nosso servidor aqui também. Vou usar a função log.Fatal(), que significa que qualquer problema que eu tenha quando eu subir o meu servidor ele vai nos apontar, vai mostrar para nós.

[08:02] E vou usar o mesmo pacote log.Fatal(http.ListenAndServe()), ou seja, escuta e sirva, escute e suba e servidor e vou passar aqui qual é a porta que eu quero subir. Vou subir no local host na porta (“:8000”, nil) mesmo, vou colocar vírgula aqui porque não temos nenhum handler para configurarmos por isso vou deixar como nil. log.Fatal(http.ListenAndServe(":8000", nil)). Não temos nenhuma configuração extra que precisaremos fazer.

package main

import (
    "fmt"
    "net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
    fmt.print(w, "Home Page")
}

func HandleRequest() {
    http.HandleFunc("/", Home)
    log.Fatal(http.ListenAndServe(":8000", nil))
}

func main() {
    fmt.Println("Iniciando o servidor Rest com Go")
    HandleResquest()
}COPIAR CÓDIGO
[08:29] Vou salvar esse nosso projeto, vou abrir aqui o nosso terminal utilizando as teclas "Ctrl + J" ou “Command + J” e vamos executar esse código: go run main.go. E ele vai dar uma mensagem: "Iniciando o servidor Rest com GO" e parou. Eu vou abrir uma aba anônima e vou colocar aqui nesse endereço “localhost:8000” e temos aqui a página Home Page.

[08:57] Até aí está tranquilo só que não queremos exibir uma página com um texto específico, queremos que quando chegue uma determinada requisição para um endereço específico queremos responder com dados, de preferência com JSON que é o padrão mais utilizado atualmente. Isso vamos aprender na sequência.

@@04
Modularizando o código

[00:00] Subimos a nossa primeira página Home só que estou preocupado com algo, observe que esse código que fizemos aqui é simples e está com 21 linhas. Vamos relembrar os passos que demos: criamos uma função Home() que quando bater uma requisição para essa função Home() vamos exibir uma mensagem retornando o ResponseWriter e a mensagem será home page.
[00:26] Depois temos alguém que só toma conta das requisições HandleRequest(), quando chegar alguém aqui nesse endpoint específico, nesse endereço específico quem vai atender vai ser o Home.

package main

import (
    "fmt"
    "net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
    fmt.print(w, "Home Page")
}

func HandleRequest() {
    http.HandleFunc("/", Home)
    log.Fatal(http.ListenAndServe(":8000", nil))
}

func main() {
    fmt.PrintIn("Iniciando o servidor Rest com Go")
    HandleResquest()
}COPIAR CÓDIGO
[00:38] Já conseguimos quebrar isso, por que? Podíamos colocar todo esse conteúdo aqui em um arquivo específico, em uma pasta específica porque a medida que o nosso projeto for aumentando temos as divisões certinhas, já sabemos exatamente onde precisamos alterar no nosso código.

[00:54] Imagine que estamos com uma página só exibindo um texto home page e estamos com 21 linhas, à medida que o nosso projeto for aumentar muito, tivermos diversos endpoints, rotas, controllers e depois modelos isso vai ficar muito difícil manter. Por isso vamos modularizar o nosso código.

[01:12] A primeira coisa que vou fazer vai ser criar uma pasta chamada routes. Vou criar aqui clicando nesse ícone aqui, um novo arquivo, um new file que vou chamar de “routes”. Aqui dentro eu vou criar um arquivo chamado “routes.go”, esse arquivo eu vou falar qual é o pacote que ele pertence que é o pacote de rotas.

[01:32] E eu vou manter nesse código de rotas esse nossa HandleResquest(). Vou tirar HandleResquest() da pasta “main.go” e vou colocá-lo aqui na pasta “routes.go” utilizando as teclas "Ctrl + C" e "Ctrl + V" e quando eu salvo ele já faz os imports certinhos para mim.

[01:45] Só que existe algo estranho aqui, ele fala que fiz os imports só que existe uma variável que está declarada e não sei o que ela é, que é essa Home. O que é Home? Ele não consegue enxergar o que está escrito aqui no nosso código “main.go”, que é a nossa função Home(). O que vamos fazer?

[02:02] Esse Home a funcionalidade dele é quando chegar uma requisição aí que lá no nosso arquivo de “routes.go” vai ser o responsável, quem vai controlar essa página vai ser essa nossa função.

[02:15] Podemos colocar isso em uma pasta chamada "controllers". Vou colocar aqui uma pasta chama “controllers” e vou criar um arquivo chamado “controllers.go”. Esse arquivo controllers eu vou colocar justamente esse nosso código Home, utilizo as teclas “Ctrl + X" para tirá-lo daqui e antes de colocá-lo aqui no “controller.go” vamos falar que package controllers, coloca aqui a nossa função, salvo e ela já está certinha.

package controllers

import (
    "fmt"
    "net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Home Page")
}COPIAR CÓDIGO
[02:43] Lá nas rotas eu vou falar que quando chegar uma requisição para barra quem vai atender vai ser alguém que está lá nos controllers e aqui eu tenho várias opções. Como a nossa função chama-se Home(), será http.HandlerFunc(“/”, controllers.Home e quando eu salvo esse projeto ele já traz o import certinho.

package routes

import (
    "log"
    "net/http"

    "github.com/guilhermeonrails/go-rest-api/controllers"
)

func HandleResquest() {
    http.HandleFunc("/", controllers.Home)
    log.Fatal(http.ListenAndServe(":8000", nil))
}COPIAR CÓDIGO
[03:13] Esse import é único, é o nome que vai estar linkado com o projeto que você iniciou lá no GoMod. Vamos supor que você colocou o nome do seu projeto só go-rest-api ele vai trazer só: go-rest-api/controllers. Se você colocou um repositório do GitHub ele vai trazer o nome completo para você também.

[03:33] Colocamos o Home() no “controllers.go”, ele fez os imports necessários do Home(), de todos os modos que precisamos para esse nosso projeto funcionar e ele já consegue identificar. No “main.go” temos um único problema, ele sabe o que é HandleResquest(), vamos falar para ele que o HandleResquest() está no nosso arquivo de rotas, é ele quem vai lidar com todas as nossas requisições.

[03:57] Ao colocar routes.HandleResquest() ele já aparece o que eu quero. Eu salvo e ele já está tudo funcionando. Ele sabe que ele precisa executar essa função HandleResquest() lá do nosso módulo de rotas.

[04:16] Vamos ver se isso está funcionando. Eu vou abrir o nosso terminal e rodar o nosso servidor, vou colocar aqui go run main.go ele vai falar: "iniciando o servidor rest com go", vou atualizar e está funcionando.

[04:33] Olha o que fizemos agora: temos um arquivo main com 13 linhas.

package main

import (
    "fmt"

    "github.com/guilhermeonrails/go-rest-api/routes"
)

func main() {
    fmt.Println("Iniciando o servidor Rest com Go")
    routes.HandleResquest()
}COPIAR CÓDIGO
[04:41] Temos um arquivo só para controllers.

package controllers

import (
    "fmt"
    "net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Home Page")
}COPIAR CÓDIGO
[04:43] E um arquivo para manter só as nossas rotas.

package routes

import (
    "log"
    "net/http"

    "github.com/guilhermeonrails/go-rest-api/controllers"
)

func HandleResquest() {
    http.HandleFunc("/", controllers.Home)
    log.Fatal(http.ListenAndServe(":8000", nil))
}COPIAR CÓDIGO
[04:45] Se surgir uma nova rota colocamos qual será o endpoint dessa rota. Só colocamos aqui o endpoint dessa nova rota e depois vincula com o controller que vamos criar e o nosso código ficará mais fácil para manter e organizar.

@@05
Respondendo Json

[00:00] Vamos evoluir a nossa aplicação e eu vou criar um estrutura para listar personalidades. Na minha cidade existem várias ruas, cujo o nome das pessoas eu não faço ideia de quem são e eu quero criar uma aplicação que cubra isso daí.
[00:14] Eu quero ficar registrando. Na minha cidade existe uma rua chamada Carmela Dutra, vou até abrir aqui para vocês. Se eu vier aqui no buscador do Google e digitar: “Avenida Carmela Dutra” ele vai mostrar que tem uma avenida Carmela Dutra também em Guarulhos. Pesquisando por “Carmela Dutra Mogi das Cruzes”, vai mostrar que em Mogi das Cruzes tem uma rua Carmela Dutra.

[00:44] Só que eu não faço ideia quem foi Carmela Dutra, vamos pesquisar: “Carmela Dutra”, quando eu pesquiso ele fala assim: "Carmela Dutra foi a primeira dama...". Foi a primeira dama do Brasil, casada com o presidente Eurico Gaspar Dutra.

Carmela Dutra: Carmela Teles Leite Dutra foi a primeira-dama do Brasil, de 31 de janeiro de 1946 até a sua morte, tendo sido a esposa de Eurico Gaspar Dutra, 16º Presidente do Brasil. Era, carinhosamente, chamada de Dona Santinha, pela sua forte religiosidade, fazendo seu marido abrir uma capelinha no Palácio Guanabara.
[01:05] Eu queria mostrar justamente isso, uma aplicação que eu mostre o nome da personalidade e mostre um pouco da história dela. Para isso vamos criar uma estrutura, o modelo de estrutura que queremos para que criemos várias personalidades da nossa cidade. Eu vou usar como exemplo nomes da minha cidade, mas você pode usar os nomes da sua cidade também e assim você descobre um pouco mais da história de onde você mora.

[01:29] A primeira que vou fazer vai ser criar uma nova pasta que eu vou chamar de models, essa pasta terá os modelos que eu vou conversar depois com o banco de dados. Vamos aprender mais para frente como isso funciona. Dentro de models eu vou criar um arquivo que eu vou chamar de “personalidades.go”. Eu preciso dizer qual é o pacote, que é o pacote de models, package models.

[01:57] E vou criar aqui uma estrutura de personalidade, eu posso digitar type Personalidade struct, abro e fecho chaves e começo a digitar a minha personalidade. type Personalidade struct { }. Outra coisa legal também para criarmos estruturas de uma forma bem rápida se você estiver usando o VS Code, é digitar apenas tys dar um "Enter" e você coloca Personalidade.

[02:31] Toda personalidade vai ter um nome e vai ter uma história, já começamos assim. O nome vai ser do tipo string, assim como a história também será uma string. Nome string, Historia string. Aqui vem uma sacada legal, para criarmos uma aplicação que seja uma Api Rest conseguimos já identificar a forma que esse nome será exibido na nossa Api Rest, como vamos serializar esse nosso campo.

[03:00] Uma das formas que temos que colocando o Nome string com json:"nome" entre crases. E vou fazer o mesmo no campo História, História strind json:"historia" também com as crases. Quando eu salvo, ele já indenta para ficar bem certinho.

[03:40] Toda personalidade minha vai ter um nome e uma história. Eu vou criar uma variável, eu vou mocar, falamos "mocar" quando não estamos de fato indo lá no banco de dados, estamos criando esses nomes na mão.

[03:51] Vou criar aqui uma variável chamada var Personalidades []Personalidades e vou falar que essa variável personalidades vai ser do tipo uma array de personalidade. Serão várias personalidades dentro dessa nossa variável. Vamos ter várias personalidade dentro da variável personalidades.

package models

type Personalidade struct {
    Nome     string `json:"nome"`
    Historia string `json:"historia"`
}

var Personalidades []PersonalidadeCOPIAR CÓDIGO
[04:09] Espero que vocês entendam. Vamos lá. Eu vou mocar algumas personalidades, como estamos lidando só com exemplo eu vou criar aqui um exemplo, vou colocar aqui mesmo. Antes de subirmos o nosso servidor, depois que iniciamos o servidor Rest eu vou colocar aqui, por exemplo, lá dos meus models eu quero trazer models.Personalidades =[]models.Personalidade {} e vou mocar alguns nomes aqui.

[04:46] Vou colocar aqui o primeiro, toda personalidade eu sei que ela tem um nome. Quando eu coloco nome dois pontos, vou deixar bem simples Nome 1 e, além do nome a personalidade também tem uma história {Nome: "Nome 1", Historia: "Historia 1"}. Vou colocar mais um aqui para termos duas personalidades como exemplo, {Nome: "Nome 2", Historia: "Historia 2"}.

package main

import (
    "fmt"

    "github.com/guilhermeonrails/go-rest-api/models"
    "github.com/guilhermeonrails/go-rest-api/routes"
)

func main() {
    models.Personalidades = []models.Personalidade{
        {Nome: "Nome 1", Historia: "Historia 1"},
        {Nome: "Nome 2", Historia: "Historia 2"},
    }

    fmt.Println("Iniciando o servidor Rest com Go")
    routes.HandleResquest()
}COPIAR CÓDIGO
[05:14] Temos agora essas duas personalidades e o que eu quero fazer na sequência é exibir essas duas personalidades em uma determinada endpoint, em um determinado endereço. Quando a pessoa digitar “localhost8000/api/personalides” eu quero exibir esses dois nomes aqui no formato JSON. Vamos ver como podemos fazer isso.

[05:32] A primeira coisa que vamos fazer para conseguir exibir essas personalidades é criando um determinado controller que é o que exibe todas as personalidades. Vou criar uma função aqui em “controller.go” que vou chamar de func TodasPersonalidades() {}. A função TodasPersonalidades() também recebe como parâmetro o ResponseWriter e Request. func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {}.

[06:02] E uma das formas que temos para conseguir colocar essa personalidade é assim: eu vou chamar do módulo json e aqui é muito importante, existe uma forma para conseguirmos responder usando o ResponseWriter que é o NewEncoder(), não é NewDecoder(), eu quero encodar uma resposta e eu quero encodar nessa resposta, json.NewEncoder(w).Encode(models.Personalides).

[06:40] Tem que colocar personalidades no plural porque se colocar só personalidade vai ser uma só. Já temos o nosso controller, é só isso que precisamos. Eu vou falar que quero ter uma resposta, quero ter um novo Enconder tipo JSON e quero encondar lá dos meus modelos todas as personalidades. Agora falta só criarmos a nossa rota.

package controllers

import (
    "encoding/json"
    "fmt"
    "net/http"

    "github.com/guilhermeonrails/go-rest-api/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Home Page")
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(models.Personalidades)
}COPIAR CÓDIGO
[06:59] Para criarmos a nossa rota determinamos o caminho, podemos colocar até esse caminho mesmo api/personalidades. Vamos ver. Vou utilizar as teclas "Shift + Ctrl" para baixo para ele copiar, vou colocar http.HandlerFunc(“/api/personalidades” e quem vai atender essa requisição serão todas as personalidades. http.HandleFunc("/api/personalidades", controllers.TodasPersonalidades).

package routes

import (
    "log"
    "net/http"

    "github.com/guilhermeonrails/go-rest-api/controllers"
)

func HandleResquest() {
    http.HandleFunc("/", controllers.Home)
    http.HandleFunc("/api/personalidades", controllers.TodasPersonalidades)
    log.Fatal(http.ListenAndServe(":8000", nil))
}COPIAR CÓDIGO
[07:26] Salvei, tenho aqui as duas personalidades mocada, vou abrir o meu servidor, utilizo as teclas "Ctrl + C" para ele parar, vou rodar mais uma vez, go run main.go. Iniciando o servidor Rest com Go, quando eu atualizo a página no “localhost:8000” temos o home page. Se eu coloco “localhost:8000/api/personalidades” temos lá o "Nome 1" e o "Nome 2" e as histórias que temos. E ele já está devolvendo um formato JSON para nós.

[07:55] Ficou legal com uma resposta JSON e vamos evoluir esse projeto ainda mais nos próximos vídeos.

@@06
Objetivo de uma API

Nesta aula, criamos uma função chamada Home destinada a página principal da aplicação e outra chamada TodasPersonalidades com intuito de exibir todas as personalidades já cadastradas, como mostra o código abaixo:
func Home(w http.ResponseWriter, r *http.Request) {
        fmt.Fprint(w, "Home Page")
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
        json.NewEncoder(w).Encode(models.Personalidades)
}COPIAR CÓDIGO
Sabendo disso, analise as afirmações abaixo e marque apenas a verdadeira.

Ambas as funções respondem uma requisição com um arquivo JSON.
 
Alternativa correta
O principal objetivo da função TodasPersonalidades não é renderizar uma página, e sim expor os dados.
 
Alternativa correta! Certo! A troca de dados entre o cliente e o servidor é extremamente popular dentro do ecossistema de API.
Alternativa correta
O principal objetivo de uma API é renderizar uma página HTML com CSS bem estilizado e bonito.

https://pt.wikipedia.org/wiki/Modelo_cliente–servidor

07
Faça como eu fiz

Chegou a hora de você seguir todos os passos realizados por mim durante esta aula. Caso já tenha feito isso, excelente. Se ainda não fez, é importante que você implemente o que foi visto nos vídeos para poder continuar com a próxima aula, que tem como pré-requisito todo o código escrito até o momento.
Caso não encontre uma solução nas perguntas feitas por alunos e alunas deste curso, para comunicar erros e tirar dúvidas de forma eficaz, clique neste link e saiba como utilizar o fórum da Alura .

https://cursos.alura.com.br/comunicando-erros-e-tirando-duvidas-em-foruns-c19

Não tem dúvidas? Que tal ajudar alguém no fórum?
: )

@@08
O que aprendemos?

Nesta aula:
Entendemos o que é uma API;
Realizamos uma requisição GET retornando um Json.
Na próxima aula:
Vamos adicionar um ID nas personalidades e conectar nossa aplicação com um banco de dados no Docker!

#### 30/09/2023

@02-Roteador, recursos por ID e Docker

@@01
Projeto da aula anterior
PRÓXIMA ATIVIDADE

Aqui você pode baixar o zip da aula 01 ou acessar os arquivos no Github!

https://github.com/alura-cursos/api-go-rest/archive/refs/heads/aula_1.zip

https://github.com/alura-cursos/api-go-rest

@@02
Novo roteador

[00:00] Conseguimos exibir tanto a personalidade 1 quanto a personalidade 2 com esses exemplos que fizemos com o nome mocado, que não estamos pegando no banco de dados e nem nada só exibindo na tela mesmo. Isso ficou legal, só que vamos evoluir um pouco mais a nossa aplicação.
[00:16] Vamos supor que eu coloque aqui em personalidades barra 1, “localhost:8000/api/personalidades/1”, eu quero exibir só ID 1. Se eu fizer isso dessa forma ele não vai encontrar ele vai abrir novamente a página Home Page porque ele não tem esse caminho ainda.

[00:30] Para conseguirmos manipular esse novo endpoint que queremos criar vamos precisar configurar esses nossos valores que fazemos. Em outras palavras, vamos precisar de roteador que vai conseguir pegar essas informações que passamos para os nossos endpoints, como ID 1 ou ID 2.

[00:47] Pensando nisso vamos usar um pacote chamado “Gorilla Mux” e vai abrir esse primeiro link do GitHub mesmo. Ele é uma rota poderosa de HTTP e URL, é justamente o que precisamos. Vamos instalar o pacote gorilla/mux e ver como conseguimos utilizá-lo.

[01:14] Aqui na documentação eu sugiro que você depois dê uma lida passo a passo, vai discorrendo para entender como funciona essa implementação do gorilla/mux. E vamos ver na prática como conseguimos usá-lo.

[01:29] A primeira coisa, para conseguirmos instalá-lo temos esse comando aqui com a configuração correta. Eu vou clicar nesse ícone para copiar ou posso selecionar tudo aqui também, utilizar as teclas "Ctrl + C" para copiar. Vamos incorporar em nosso projeto o gorilla/mux.

[01:45] Para isso eu vou parar o nosso servidor, limpo a tela utilizando as teclas "Ctrl + L", vou utilizar as teclas "Ctrl + C" e "Ctrl + V". go get -u github.com/gorilla/mux, dou um "Enter" e vai instalar esse pacote. Se formos aqui em nosso “go.mod” repare que ele falou que temos o gorilla/mux aqui também no nosso pacote como dependência desse projeto. Isso ficou ótimo, é justamente o que queríamos.

[02:20] Para conseguirmos utilizar agora o gorilla/mux é muito simples, vamos lá na nossa rota e vou criar uma instância dele que vou chamar de r, r := mux.NewRouter(). Ou seja, ele vai criar uma nova rota, um novo mapeamento de rotas.

[03:00] A única coisa que vou fazer vai ser colocar essa instancia que temos aqui na frente: r.http.HandleFunc(), r.http.HandleFunc() e aqui quando subimos o nosso servidor no lugar de deixar o new eu vou passar a nossa instancia. Aqui não vamos precisar do r.http.HandleFunc, vamos usar apenas o r.HandleFunc() do gorilla/mux.

package routes

import (
    "log"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/guilhermeonrails/go-rest-api/controllers"
)

func HandleResquest() {
    r := mux.NewRouter()
    r.HandleFunc("/", controllers.Home)
    r.HandleFunc("/api/personalidades", controllers.TodasPersonalidades)
    log.Fatal(http.ListenAndServe(":8000", r))
}COPIAR CÓDIGO
[03:30] Vamos executar mais uma vez o nosso projeto e limpar o terminal. go run main.go. Se viermos aqui vamos que esse daqui esse daqui está carregando certinho e se eu coloco lá na nossa pasta no “localhost:8000” abre a página home certinho. Tudo está funcionando como queríamos.

[03:50] No lugar de usar o default, o mux default que vem do global estamos utilizando o mux com NewRouter() do gorilla/mux. Lembra que eu falei, agora conseguimos manipular melhor as nossas rotas e as nossas requisições.

[04:05] Na sequência, veremos como fazemos para conseguirmos criar um endpoint para visualizar o ID passado. Se colocamos o id1 aqui queremos exibir só a personalidade 1 ou só a personalidade 2 com id2 e assim por diante.

@@07
Exibindo uma personalidade
PRÓXIMA ATIVIDADE

Uma pessoa não conseguiu exibir uma personalidade por ID e resolveu pedir a sua ajuda. Veja o código que a pessoa compartilhou:
models.go
package models

type Personalidade struct {
        Id       int    `json:"id"`
        Nome     string `json:"nome"`
        Historia string `json:"historia"`
}

var Personalidades []PersonalidadeCOPIAR CÓDIGO
routes.go
func HandleResquest() {
        r := mux.NewRouter()
        r.HandleFunc("/", controllers.Home)
        r.HandleFunc("/api/personalidades", controllers.TodasPersonalidades).Methods("Get")
        r.HandleFunc("/api/personalidades/{ID}", controllers.RetornaUmaPersonalidade).Methods("Get")
        log.Fatal(http.ListenAndServe(":8000", r))
}COPIAR CÓDIGO
controllers.go
func RetornaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        id := vars["id"]

        for _, personalidade := range models.Personalidades {
                if strconv.Itoa(personalidade.Id) == id {
                        json.NewEncoder(w).Encode(personalidade)
                }
        }
}COPIAR CÓDIGO
Analisando os trechos de código acima, podemos afirmar que:

Na função HandleResquest do arquivo routes.go, o método da rota deveria ser Post e não Get.
 
Incorreto. O método GET solicita a representação de um recurso específico. Requisições utilizando o método GET devem retornar apenas um recurso, seja ele com todas as personalidades ou apenas uma. Já o método POST é utilizado para submeter uma entidade a um recurso específico, frequentemente causando uma mudança no estado do recurso ou efeitos colaterais no servidor. Ou seja, usamos o método POST para criar um novo recurso, no nosso caso, uma nova personalidade.
Alternativa correta
O erro se encontra na comparação dos IDs, especificamente neste trecho de código if strconv.Itoa(personalidade.Id) == id, onde o personalidade.Id deveria ter a letra i em minúsculo para a comparação ter efeito.
 
Alternativa incorreta. Erro não se encontra neste trecho de código. Toda variável pública que importamos de outro pacote, deve ser identificada pela primeira letra maiúscula.
Alternativa correta
A função do arquivo controllers.go que retorna uma personalidade, não encontrará o ID das variáveis de rota que podem ser recuperadas chamando mux.Vars.
 
Alternativa correta! Isso! Observe que o parâmetro da rota /api/personalidades/{ID} expõe uma variável chamada ID, com todas as letras maiúsculas e a função que retorna todas as personalidades busca uma variável chamada id, com letras minúsculas.

@@08
Faça como eu fiz
PRÓXIMA ATIVIDADE

Chegou a hora de você seguir todos os passos realizados por mim durante esta aula. Caso já tenha feito isso, excelente. Se ainda não fez, é importante que você implemente o que foi visto no vídeo para poder continuar com a próxima aula, que tem como pré-requisito todo o código escrito até o momento.
Caso não encontre uma solução nas perguntas feitas por alunos e alunas deste curso, para comunicar erros e tirar dúvidas de forma eficaz, clique neste link e saiba como utilizar o fórum da Alura .

https://cursos.alura.com.br/comunicando-erros-e-tirando-duvidas-em-foruns-c19

Não tem dúvidas? Que tal ajudar alguém no fórum?

: )

@@09
O que aprendemos?
PRÓXIMA ATIVIDADE

Nesta aula:
Adicionamos o Gorilla Mux como novo roteador da nossa aplicação;
Retornamos um único recurso através do id;
Criamos uma imagem do banco de dados Postgres com Docker e executamos um script SQL que adicionava alguns registros em nosso banco de dados.
Na próxima aula:
Vamos aprender como conectar nosso projeto com banco de dados e exibir as informações armazenadas utilizando um ORM!

https://github.com/alura-cursos/api-go-rest/blob/aula_2/routes/routes.go

https://github.com/alura-cursos/api-go-rest/blob/aula_2/controllers/controllers.go

https://github.com/alura-cursos/api-go-rest/blob/aula_2/migration/docker-database-initial.sql

#### 02/10/2023

@03-Conexão com banco e exibindo os dados

@@01
Projeto da aula anterior
PRÓXIMA ATIVIDADE

Aqui você pode baixar o zip da aula 02 ou acessar os arquivos no Github!

https://github.com/alura-cursos/api-go-rest/archive/refs/heads/aula_2.zip

https://github.com/alura-cursos/api-go-rest/tree/aula_2

@@02
Sobre o GORM

[00:00] O próximo passo é conseguirmos conectar com o banco de dados e exibir as personalidades que temos registradas lá. Para isso vamos utilizar um ORM do GO para conseguirmos pedir as informações para retornar um registro do banco de dados ou mostrar todas as personalidades, criar uma personalidade de uma forma mais simples.
[00:22] No curso anterior não usamos o ORM então escrevemos SQL a mão. Nesse curso vamos utilizar o Gorm, vamos digitar aqui no buscador do Google "gorm orm" e ele vai trazer essa opção aqui "GORM - A fantástica biblioteca de ORM para Golang".

[00:40] Para instalarmos o gorm na nossa aplicação vamos rodar esse comando aqui go get -u gorm.io/gorm. Vou rodar aqui: go get -u gorm.io/gorm, dou um "Enter" e ele está trazendo o Gorm para a nossa aplicação.

[00:58] Agora vamos ver na documentação como é que fazemos para conseguir conectar com o banco de dados. Vou colocar aqui "Connecting to a Database" e temos a opção para ele conectar com MySQL, PostgreSQL, SQLite e o Server. Vamos usar o postgres.

[01:15] Indo um pouco para baixo repare que ele tem sempre um driver de conexão, ele tem o driver de conexão e ele fala para sempre inicializarmos uma variável apontando para Gorm.db, quando queremos já tem um banco de dados e queremos conectar com esse banco de dados. É justamente o que queremos.

[01:42] Já existe o banco de dados e queremos fazer a conexão com esse banco de dados. Para isso eu vou criar uma pasta chamada “database” e dentro dela um arquivo chamado “db.go”. Esse arquivo “db.go” vai ser responsável por conectar com o banco de dados.

[01:59] Vou falar qual é o pacote database, package database e vamos precisar de duas variáveis que já vou criar aqui. A primeira variável vou chamar de DB, ambos maiúsculas, que vai apontar para *gorm. Aqui temos um ponto importante, só de eu digitar gorm aparece aqui "gorm.io/gorm" e o "gorm.io.jinzuh" desses dois.

[02:22] Vamos usar a primeira opção, "gorm.io/gorm". Dou um "Enter" nesse aqui e repare que ele já fez o import do gorm. Tem várias opções aqui vamos colocar ponto DB, seguindo a documentação. var { DB *gorm.DB }. Ele mostra aqui que se já temos um banco de dados, quando já temos uma conexão existente usamos um apontador para gorm DB.

[02:45] Outra coisa que vamos precisar é do driver do postgres também. Eu vou vir aqui e vou colocar go get gorm.io/driver/postgres no terminal, dou um "Enter" e ele vai nos trazer a versão do postgres para utilizarmos.

[03:01] Outra variável que vamos precisar também é uma variável de erro, err error. Agora eu vou criar uma função que vou chamar de func ConectaComBancodeDados() e essa função vai ser a seguinte, vou seguindo a documentação mesmo.

[03:23] Existe uma linha, uma string de conexão, vou utilizar as teclas "Ctrl + C" para copiar essa linha aqui e vou colar aqui utilizando as teclas "Ctrl + V". Vou mudar o nome, vou chamar de string de conexão. stringDeConexao := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disableTimeZone=Asia/Shanghai".

[03:39] Vamos ver como essa linha está. Temos o host que é igual a localhost, está certo. O user igual tem que ser igual a root e não gorm. Lembrando que esse nome root estamos trazendo lá do nosso “docker-compose”, user=root, password=root e o banco de dados root também. Com base nisso vamos colocar user=root, o password=root também e o banco de dados root.

[04:03] A porta que estamos utilizando é a 5432, o sslmode vou deixar desativado e vou tirar aqui o TimeZone. stringDeConexao := "host=localhost user=root password=root dbname=root. port=5432 sslmode=disable". Salvei essa string e agora vou precisar conectar com o Gorm, conectar essa string que fizemos com o banco de dados.

[04:21] Vou colocar essas duas variáveis que criamos tanto o DB quanto o erro e vou colocar apenas o igual, não tem nesse daqui os dois pontos porque já inicializamos essas nossas variáveis indicando o tipo delas, DB, err =.

[04:35] Vou falar assim DB, err = gorm.Open e aqui dentro vou pedir para o postgres abrir para mim, postgres. e temos algumas opções aqui, vou colocar a opção Open() e ele vai abrir também. O postgres vai ser carregado com base na nossa stringDeConexao. DB, err = gorm.Open(postgres.Open(stringDeConexao).

[04:58] Eu salvo, está tudo certo e a única coisa que precisamos fazer é verificar se temos algum erro. Se tivermos algum erro significa que não conseguimos conectar com o banco de dados, eu vou colocar aqui log.Panic() para visualizarmos a mensagem erro ao conectar com banco de dados. log.Panic("Erro ao conectar com banco de dados").

package database

import (
    "log"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var (
    DB  *gorm.DB
    err error
)

func ConectaComBancoDeDados() {
    stringDeConexao := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
    DB, err = gorm.Open(postgres.Open(stringDeConexao))
    if err != nil {
        log.Panic("Erro ao conectar com banco de dados")
    }
}COPIAR CÓDIGO
[05:20] Salvando essa linha parece que está tudo ok e eu vou pedir para conectar com o banco de dados lá na nossa função main(). Na nossa função main criamos os nossos modelos, iniciamos o servidor. Eu vou deixar esses modelos por enquanto e vou colocar lá de database, vou dar um "Enter" aqui para ele importar o database e ele já nos trouxe guilherme/database que acabamos de criar e eu quero conectar com o banco de dados, database.ConectaComBancoDeDados.

package main

import (
    "fmt"

    "github.com/guilhermeonrails/go-rest-api/database"
    "github.com/guilhermeonrails/go-rest-api/models"
    "github.com/guilhermeonrails/go-rest-api/routes"
)

func main() {
    models.Personalidades = []models.Personalidade{
        {Id: 1, Nome: "Nome 1", Historia: "Historia 1"},
        {Id: 2, Nome: "Nome 2", Historia: "Historia 2"},
    }
    database.ConectaComBancoDeDados()
    fmt.Println("Iniciando o servidor Rest com Go")
    routes.HandleResquest()
}COPIAR CÓDIGO
[05:49] Salvando eu vou limpar o meu terminal e vou rodar o nosso projeto go run main.go, dei um "Enter" e não recebemos nenhuma mensagem. Repare que conectamos com o banco de dados e depois iniciamos o nosso servidor. Estamos conectados com o banco de dados que estamos rodando no Docker.

[06:13] O próximo passo é no lugar de exibirmos as personalidades que criamos aqui apenas de teste, eu quero exibir de fato as personalidades que estão no banco de dados.

@@03
Controller e GORM

[00:00] Vamos alterar o nosso controller para que ele exiba as personalidades que temos cadastrada no banco de dados já que já estamos conectados.
[00:10] A primeira coisa que vamos fazer vai ser alterar o nosso controller, vou vir aqui "controller > controllers.go" e repare o que fazemos, temos todas as personalidades que mocamos lá no nosso main() para gerar esse nosso exemplo. Não é o que queremos, queremos esses dados que estão aqui.

[00:32] Para exibir esses dados a primeira coisa que vamos fazer vai ser criar uma variável do tipo lista de personalidades aqui dentro. Vamos fazer passo a passo. Eu vou criar uma nova variável que chamar de p que vai ser do tipo Personalidade, uma lista que vem lá de models.Personalidades. var p []models.Personalidade.

[00:52] Isso significa que essa minha variável p vai ser uma lista com base em todas as minhas personalidades. Agora que eu já tenho essa lista de personalidades eu vou pedir para o meu database que é a nossa instância do GORM, encontrar todas as personalidades que ele tem e para ele encontrar vamos editar Find() que vai encontrar todas as personalidades. database.DB.Find().

[01:20] Para ele conseguir saber o que ele está buscando vamos passar o endereço de memória de p. O p é uma lista, o p é um array de personalidades e ele vai encontrar todas essas listas de personalidades passando o endereço de memória de p. database.DB.Find(&p).

[01:42] Depois disso, no lugar de visualizarmos o models.personalidade que temos lá do main() vamos visualizar o p. json.NewEncoder(w).Encode(p). Eu vou testar isso aqui para ver porque parece que foi muito simples isso daqui. Lembra do curso passado que tínhamos que criar para encontrar todos e escrevia select asterisco from personalidades e tudo aquilo, não precisamos. Com uma linha de código conseguimos conectar com o banco de dados e trazer todos eles.

package controllers

import (
    "encoding/json"
    "fmt"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/guilhermeonrails/go-rest-api/database"
    "github.com/guilhermeonrails/go-rest-api/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Home Page")
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
    var p []models.Personalidade
    database.DB.Find(&p)
    json.NewEncoder(w).Encode(p)
}

func RetornaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    for _, personalidade := range models.Personalidades {
        if strconv.Itoa(personalidade.Id) == id {
            json.NewEncoder(w).Encode(personalidade)
        }
    }
}COPIAR CÓDIGO
[02:11] Vamos rodar go run main.go, ele conectou com o banco de dados e iniciou um servidor. Lembrando que o banco de dados precisa estar rodando no Docker, se o Docker não estiver ativo, não estiver funcionando não conseguimos conectar com o banco de dados.

[02:36] O Docker está funcionando, iniciamos o servidor e no banco de dados temos Deodato e a Carmela Dutra e aqui temos esses nomes de exemplo. Vou tirar esses dois, dar um “Enter” e olha lá, temos o Deodato e temos também a Carmela Dutra, aparecem esses dois dados.

[02:58] Dessa forma, conseguimos trazer as informações que estão lá no banco de dados para a nossa API com GO. Isso fizemos através do GORM com essa linha de comando. Na sequência vamos alterar as nossas funcionalidades, principalmente essa que retorna uma determinada personalidade através do ID utilizando o GORM.

@@04
Buscando um recurso

[00:00] O nosso projeto está até com um comportamento interessante. Quando pedimos para exibir todas as personalidades ele de fato exibe personalidades que fazem sentido para mim, o nome da rua, a história. Só que quando eu quero ver só a história do personagem 1 ele mostra Nome 1, Id: 1. Ou seja, ele está pegando os dados que mocamos lá no nosso main(), não faz muito sentido.
[00:21] Eu quero colocar o 1 e exibir de fato a personalidade que está lá no nosso banco de dados com Id: 1. Para isso vamos pedir para o nosso ORM fazer isso.

[00:32] Pegamos o valor dessa variável, temos aqui o valor dessa variável por Id, id := vars("id"), até aqui está tudo certo. O que vamos fazer agora? Nós criamos uma lista de personalidades para conseguirmos visualizar todas elas, nesse nosso caso vamos criar apenas uma personalidade, var p models.Personalidade. Vai ser apenas um, esse vai ser o tipo dela.

[00:59] Eu vou tirar todas essas linhas aqui, tudo isso que fizemos eu vou fazer da seguinte forma: database.DB.First. O .First vai trazer a primeira referência que ele encontrar e o Id é algo único.

[01:19] Aqui temos um ponto importante: queremos o endereço de memória de p, vou passar da mesma forma que fizemos para o outro, passar o endereço de memória de &p que é essa nossa variável aqui. Eu sei que você está buscando uma personalidade, mas qual personalidade específica? Eu preciso saber a outra informação que vai ser o id, eu vou passar aqui id. database.DB.First(&p, id)

[01:43] Dessa forma eu vou falar assim: json.NewEncoder(w).Encode(p). Para o nosso código ficar ainda mais claro eu vou alterar, no lugar de p eu vou deixar personalidade para todos. var personalidade models.Personalidade, database.DB.First(&personalidade, id), json.NewEncoder(w).Encode(personalidade).

[02:14] Vou salvar. Parece que está tudo funcionando, não precisamos fazer nenhuma alteração em nenhuma outra parte da nossa aplicação. A não ser fechar o nosso servidor GO e subir o nosso servidor GO para ver se está tudo ok.

package controllers

import (
    "encoding/json"
    "fmt"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/guilhermeonrails/go-rest-api/database"
    "github.com/guilhermeonrails/go-rest-api/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Home Page")
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
    var p []models.Personalidade
    database.DB.Find(&p)
    json.NewEncoder(w).Encode(p)
}

func RetornaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    var personalidade models.Personalidade
    database.DB.First(&personalidade, id)
    json.NewEncoder(w).Encode(personalidade)
}COPIAR CÓDIGO
[02:28] Vou vir aqui na nossa aplicação no navegador, "/personalidades/1" quando eu dou um "Enter" agora sim, tenho o doutor Deodato. Quando eu coloco aqui a "/personalidade/2" ele mostra a Carmela Dutra, se eu coloco uma personalidade que não tenha, "/50" por exemplo, ele mostra tudo vazio.

[02:43] Dessa forma, aos poucos vamos entendendo como funciona esse nosso ORM. O legal é que ele vai mostrando o que ele está encontrando, quais são os scripts que ele está executando para conseguir exibir todas as personalidades. Isso é muito bacana, é bem legal também para visualizarmos.

[03:04] Vimos que usando o Find() buscamos todas as personalidades passando o endereço de memória da estrutura que temos, que queremos exibir e para retornarmos apenas uma utilizamos o DB.First para ele voltar passando o endereço de memória da personalidade, da pessoa que estamos buscando e o Id.

@@05
Go mod
PRÓXIMA ATIVIDADE

Conectamos nossa aplicação com banco de dados Postgres e adicionamos alguns módulos nesta aula. Podemos ver todos os módulos e pacotes usados até aqui no arquivo go.mod:
module github.com/guilhermeonrails/go-rest-api

go 1.15

require (
        github.com/gorilla/mux v1.8.0
        gorm.io/driver/postgres v1.1.2
        gorm.io/gorm v1.21.16 // indirect
)COPIAR CÓDIGO
Sabendo disso, analise as afirmações abaixo e marque a que está correta em relação aos módulos usados até aqui.

A conexão do banco de dados com o GORM deve ser feita exclusivamente com o Docker. Caso contrário, a conexão não é possível.
 
Alternativa correta
O único banco de dados que o Gorm aceita é o Postgres.
 
Alternativa correta
O pacote gorilla/mux implementa um roteador de requisições e respostas para corresponder às solicitações de entrada ao seu respectivo manipulador ou handler.
 
Alternativa correta! Certo! Ele encaminha o handler específico com base no método da requisição. Podemos ver esse exemplo no arquivo routes.go.

https://github.com/alura-cursos/api-go-rest/blob/aula_2/routes/routes.go

@@06
Faça como eu fiz
PRÓXIMA ATIVIDADE

Chegou a hora de você seguir todos os passos realizados por mim durante esta aula. Caso já tenha feito isso, excelente. Se ainda não fez, é importante que você implemente o que foi visto no vídeo para poder continuar com a próxima aula, que tem como pré-requisito todo o código escrito até o momento.
Caso não encontre uma solução nas perguntas feitas por alunos e alunas deste curso, para comunicar erros e tirar dúvidas de forma eficaz, clique neste link e saiba como utilizar o fórum da Alura .

https://cursos.alura.com.br/comunicando-erros-e-tirando-duvidas-em-foruns-c19

Não tem dúvidas? Que tal ajudar alguém no fórum?
: )

@@07
O que aprendemos?
PRÓXIMA ATIVIDADE

Nesta aula:
Instalamos o Gorm;
Realizamos a conexão da API com banco de dados;
Alteramos as funções do controller para exibir as informações do banco de dados.
Na próxima aula:
Vamos aprender como criar, editar e deletar uma personalidade!