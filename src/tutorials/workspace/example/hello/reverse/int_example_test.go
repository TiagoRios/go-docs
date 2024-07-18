// se estivesse no mesmo pacote não precisaria importar
package reverse_test

import (
	"fmt"

	"golang.org/x/example/hello/reverse"
)

func ExampleInt() {
	fmt.Println(reverse.Int(123))
	// Output: 321
}
