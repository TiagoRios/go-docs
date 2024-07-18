package methodsinterfaces

import (
	"fmt"
	"math"
)

// ---------- interfaces ----------

type Abser interface {
	Absolute() float64
}

// A type implements an interface by implementing its methods.
// There is no explicit declaration of intent, no "implements" keyword.
func (v *Vertex) Absolute() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// ---------- values Interfaces ----------

type MinhaInterface interface {
	MeuMetodo()
}

type MeuTipo struct {
	Mensagem string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t *MeuTipo) MeuMetodo() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.Mensagem)
}

// This type implements the interface too
type MeuTipoFloat float64

func (f MeuTipoFloat) MeuMetodo() {
	fmt.Println(f)
}

// Somente armazena valores do tipo dela.
// func describe(i MinhaInterface) {
// 	fmt.Printf("(%v, %T)\n", i, i)
// }

// An empty interface may hold values of any type.
// (Every type implements at least zero methods.)
// Empty interfaces are used by code that handles values of
// unknown type. For example, fmt.Print takes any number of
// arguments of type interface{}.
func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

// ---------- Type Assertions ----------

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}
