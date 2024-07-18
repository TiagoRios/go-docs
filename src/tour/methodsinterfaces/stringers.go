package methodsinterfaces

import "fmt"

// definiido pelo pacote fmt
// type Stringer interface {
// String() string
// }

// A Stringer is a type that can describe itself as a string.
// The fmt package (and many others) look for this interface
// to print values.

type Person struct {
	Name string
	Age  int
}

// implementando a interface Stringer
// funciona como se fosse um toString() do java
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

// ---------- Exercicio stringer ----------

// Make the IPAddr type implement fmt.Stringer to print the
// address as a dotted quad.
// For instance, IPAddr{1, 2, 3, 4} should print as "1.2.3.4".

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.

func (ip IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
}
