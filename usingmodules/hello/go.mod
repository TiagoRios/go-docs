module example.com/hello

go 1.21.3

// Posso utilizar o repositório local desta maneira.
replace example.com/greetings => ../greetings

// gerou um pseudo número de versão.
require example.com/greetings v0.0.0-00010101000000-000000000000
