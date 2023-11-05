# Aprendendo a linguagem GO

Este repositório contém os códigos dos tutoriais disponíveis na documentação da linguagem GO (GoLang) 

## Iniciando

Instale o GO;

Crie um diretório para o projeto;

Crie um arquivo com a extensão ".go" dentro do diretório.

```go
// hello.go

package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

Criar o modulo:

```bash
$ go mod init example/hello
```

Executar o código:

```bash
$ go run .
```

Resolver as dependências:

```bash
$ go mod tidy
go: finding module for package rsc.io/quote
go: found rsc.io/quote in rsc.io/quote v1.5.2
```
<p>
O comando acima busca os modulos para os pacotes externos que foram importados em seus arquivos ".go".
</p> 

Para usar seu próprio modulo local quando ele ainda não estiver publicado em um repositório remoto, utilize:

```bash
$ go mod edit -replace example.com/greetings=../greetings
```

Após o comando acima, meu arquivo go.mod foi modificado: 

```
// cut
replace example.com/greetings => ../greetings
```


Atualize as dependências novamente:

```bash
$ go mod tidy
```
Depois de atualizar, meu arquivo go.mod foi modificado:

```
// cut
require example.com/greetings v0.0.0-00010101000000-000000000000
```

## [Aprenda GO](https://go.dev/doc/) hoje !!!