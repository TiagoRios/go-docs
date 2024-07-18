package methodsinterfaces

import "fmt"

func ExampleErrNegativeSqrt_deveLancarErro() {
	if v, err := Sqrt(-4); err != nil {
		fmt.Println(err)
		fmt.Println("valor verificado:", v)

		// Output:
		// cannot Sqrt negative number: -4
		// valor verificado: -4
	}
}

func ExampleErrNegativeSqrt_naoDeveLancarErro() {
	if v, err := Sqrt(16); err == nil {
		fmt.Println("Sqrt igual:", v)

		// Output:
		// Sqrt igual: 4
	}
}
