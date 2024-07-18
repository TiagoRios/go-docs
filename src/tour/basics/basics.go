package basics

import (
	"fmt"
	"math"
	"math/rand"
)

// declaração em nível de pacote:
var one int = 1
var (
	two, three        = 2, 3 // tipo implicito
	numeroNaoIniciado int    // tipo explicito
)

func ImprimeTipoValorConteudoDaString(texto string) {
	fmt.Printf(" Type: %T;\n Value: %v;\n conteudo: %q\n", texto, texto, texto)
}

// conversão de tipos
func Hipotenusa(b, c int) {
	// a²= b²+c²
	var a float64 = math.Sqrt(float64(b*b + c*c))
	fmt.Printf("A hipotenusa de um triângulo retângulo com catetos %v e %v é: %v\n", b, c, uint(a))
}

// func Somar(x, y int) int
func Somar(x int, y int) int {
	return x + y
}

// pode retornar qualquer número de resultados
func Swap(x, y string) (string, string) {
	return y, x
}

// valores de retorno nomeado
func Split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return // naked return - use apenas em funções pequenas
}

func InteiroAleatorio(limite int) int {
	return rand.Intn(limite)
}
