package basics

import (
	"fmt"
	"math"
	"strings"
)

// ---------- ponteiros ----------

func UsandoPonteiros() {
	i, j := 42, 2701

	var ponteiro = &i // point to i

	fmt.Println("i pelo ponteiro", *ponteiro)      // read i through the pointer
	*ponteiro = 21                                 // set i through the pointer
	fmt.Println("i modificado usando ponteiro", i) // see the new value of i

	ponteiro = &j                             // point to j
	fmt.Println("j pelo ponteiro", *ponteiro) // read j through the pointer

	*ponteiro = *ponteiro / 37                     // divide j through the pointer
	fmt.Println("j modificado usando ponteiro", j) // see the new value of j
}

// ---------- structs ----------

type Vertex struct {
	EixoX int
	EixoY int
}

func UsandoStructs() {
	vertice := Vertex{1, 2}
	g := &vertice

	(*g).EixoX = 47
	g.EixoY = 3 // permite não usar desreferencia

	fmt.Println("Vertice X:", vertice.EixoX)
	fmt.Println("Vertice Y:", vertice.EixoY)
}

func StructsLiterals() {
	var (
		v1 = Vertex{1, 2}     // has type Vertex
		v2 = Vertex{EixoX: 1} // eixoY:0 is implicit
		v3 = Vertex{}         // eixoX:0 and Y:0
		p  = &Vertex{1, 2}    // has type *Vertex
	)

	fmt.Println(v1, v2, v3, p)
}

// ---------- arrays ----------

func UsandoArrays() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World asdf"
	fmt.Println(a[0], a[1])
	fmt.Println(a) // estranho - array não tem separação com vírgulas

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

// ---------- slices ----------

func UsandoSlice() {
	primes := [6]int{2, 3, 5, 7, 11, 13} // array literal
	// primes :=  []int{2, 3, 5, 7, 11, 13} // slice literal

	var meuSlice []int = primes[1:4] // 3 5 7 len3, cap5
	fmt.Println(meuSlice)
	fmt.Println("length", len(meuSlice))
	fmt.Println("capacidade", cap(meuSlice))

	var s []int // nil
	fmt.Println(s, len(s), cap(s))

	// these slice expressions are equivalent:
	// a[0:10]
	// a[:10]
	// a[0:] cópia
	// a[:] cópia
}

func SliceReferenceArray() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	// A mudança afeta outros slices e o array original.
	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}

func SliceMake() {
	a := make([]int, 5) //len 5
	printSlice("a", a)

	b := make([]int, 0, 5) //len 0, cap 5
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)
}

func printSlice2(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func SliceOfSlice() {
	// Create a tic-tac-toe board.
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
		// []string{"_", "_", "_"}, antes problema com lint
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func SliceAppend() {
	var s []int
	printSlice2(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice2(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice2(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4, 5)
	printSlice2(s)
}

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func UsandoRange() {
	// range retorna 2 valores
	// for _, valor := range pow { PULANDO index - range continuado
	// for index, _ := range pow { PULANDO valor - range continuado
	for index, valor := range pow {
		fmt.Printf("2**%d = %d\n", index, valor)
	}
}

type Coordenada struct {
	Lat, Long float64
}

var m map[string]Coordenada

// m = make(map[string]Coordenada) exemplo de atribuição

func UsandoMaps() {
	// chave tipo string, valor tipo Coordenada
	m = make(map[string]Coordenada)

	m["Bell Labs"] = Coordenada{40.68433, -74.39967}

	fmt.Println(m["Bell Labs"])

	// verifica se existe a chave no mapa
	// valor, ok := m["Bell Labs"]
	var valor, ok = m["Bell Labs"]

	fmt.Println("The value:", valor.Long, "Present?", ok)
}

func UsandoMapLiteral() {
	var m = map[string]Coordenada{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408}, // Map literals continued
		"Outro":     {11.11111, -22.22222},
	}

	fmt.Println(m["Google"].Long)
}

func UsandoMapMutacao() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	// var v, ok = m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}

// recebe uma função que aceita 2 argumentos e retorna a função
// com argumentos preenchidos.
func Compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

// Utilizando a função Compute
func FunctionValue() {
	// esta função satisfaz a função Compute
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	soma := func(x, y float64) float64 {
		return x + y
	}

	// função completa inserida diretamente.
	resultadoCompute := Compute(func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	})

	fmt.Println(hypot(5, 12)) //13

	// apenas passei a definição da função
	fmt.Println(Compute(hypot))    // 5
	fmt.Println(resultadoCompute)  // 5
	fmt.Println(Compute(math.Pow)) // 81
	fmt.Println(Compute(soma))     // 7
}

// closure
// func adder() (func(int) int) {
func adder() func(int) int {
	sum := 0

	return func(x int) int {
		sum += x
		return sum
	}
}

func FunctionClosure() {
	pos, neg := adder(), adder()

	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

// retorna uma função que pode ser executads
// prev e next continuam acessiveis e são atualizadas
// a cada chamada da função retornada. (Closure)
func fibonacci() func() int {
	prev := 0
	next := 1

	return func() int {
		atual := next + prev
		prev = next
		next = atual

		return prev
	}
}

func FibonacciClosure(quantidade int) {
	numeroFibonacci := fibonacci()

	for i := 0; i < quantidade; i++ {
		fmt.Println(numeroFibonacci())
	}
}

// ---------- exercicos slices ----------

// exemplo da documentação
// funciona apenas no tour do golang
// func main() {
// 	f := func(dx, dy int) [][]uint8 {
// 		ss := make([][]uint8, dy)
// 		for y := 0; y < dy; y++ {
// 			s := make([]uint8, dx)
// 			for x := 0; x < dx; x++ {
// 				s[x] = uint8((x + y) / 2)
// 			}
// 			ss[y] = s
// 		}
// 		return ss
// 	}

// 	pic.Show(f)

// }

// func Pic(dx, dy int) (picture [][]uint8) {
// 	picture = make([][]uint8, dy)
// 	for y := range picture {
// 		picture[y] = make([]uint8, dx)
// 		for x := range picture[y] {
// 			if y % 2 == 1 {
// 				picture[y][x] = uint8((y+x*2)^(y-x*2))
// 			} else {
// 				picture[y][x] = uint8((x+y*2)^(x-y*2))
// 			}
// 		}
// 	}
// 	return picture
// }
