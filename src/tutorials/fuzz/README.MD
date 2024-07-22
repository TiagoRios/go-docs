# Fuzzing em GO (Testes de difusão)

Executar todos os testes

```bash
$ go test .
```

Executar teste separadamente:

```bash
$ go test -run=FuzzNomeDoTeste
$ go test -run=TestNomeDoTeste
```

Execute FuzzNomeDoTeste com fuzzing, para ver se alguma entrada de string gerada aleatoriamente causará uma falha.

```bash
$ go test -fuzz=Fuzz
ou
$ go test -fuzz=FuzzNomeDoTeste
```

Ocorreu uma falha durante o fuzzing, e a entrada que causou o problema é gravada em um arquivo corpus semente que será executado na próxima vez que go testfor chamado, mesmo sem o sinalizador, -fuzz.

O arquivo terá esse formato:

```go
go test fuzz v1
string("Ҝ")
```

A string foi gerada aleatoriamente pelo teste de difusão.
