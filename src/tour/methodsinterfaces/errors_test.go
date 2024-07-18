package methodsinterfaces

import (
	"testing"
)

func TestMyError_deveLancarErro(t *testing.T) {

	// when
	err := run()

	// then - err igual a nil NÃO lançou erro = FAIL
	if err == nil {
		t.Error(err)
	}
}

func TestErrNegativeSqrt_deveLancarErro(t *testing.T) {
	if _, err := Sqrt(-9); err == nil {
		t.Error(err)
	}
}
