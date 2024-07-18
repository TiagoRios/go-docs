package methodsinterfaces

import (
	"math"
)

type Vertex struct {
	X, Y float64
}

// A method is a function with a special receiver argument.
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Here's Abs written as a regular function with no change in functionality.
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// ---------- Methods continued ----------

type MyFloat float64

// You can declare a method on non-struct types, too.
//
// You can only declare a method with a receiver whose type
// is defined in the same package as the method. You cannot
// declare a method with a receiver whose type is defined in
// another package (which includes the built-in types such as int).
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// ---------- Pointer receivers ----------

// Methods with pointer receivers can modify the value to
// which the receiver points (as Scale does here). Since
// methods often need to modify their receiver, pointer
// receivers are more common than value receivers.
// observação:
// In general, all methods on a given type should have either
// value or pointer receivers, but not a mixture of both.
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}
