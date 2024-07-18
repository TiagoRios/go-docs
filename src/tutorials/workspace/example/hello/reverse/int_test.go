package reverse

import (
	"fmt"
	"testing"
)

func TestInt_DeveReverterInteiro(t *testing.T) {

	input, output := 123, 321

	result := Int(input)

	if result != output {
		t.Errorf("reverse.Int(%q) == %q, esperava %q", input, result, output)
	}
}

// esta no mesmo pacote então chamo sem referenciar arquivo externo
func ExampleInt() {
	fmt.Println(Int(12345))
	// Output: 54321
}
