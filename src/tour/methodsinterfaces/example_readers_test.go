package methodsinterfaces

import (
	"io"
	"os"
	"strings"
)

func ExampleLeitor() {
	Leitor("Hello, Reader!", 8)
	// Output:
	// bytes lidos = 8
	// slice bytes lidos = [72 101 108 108 111 44 32 82]
	// conteudo bytes lido = "Hello, R"
	// err = <nil>
	// bytes lidos = 6
	// slice bytes lidos = [101 97 100 101 114 33 32 82]
	// conteudo bytes lido = "eader!"
	// err = <nil>
}
func ExampleLeitor_outrosDados() {
	Leitor("abc", 5)
	// Output:
	// bytes lidos = 3
	// slice bytes lidos = [97 98 99 0 0]
	// conteudo bytes lido = "abc"
	// err = <nil>
}

func ExampleRot13Reader() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := Rot13Reader{s}
	io.Copy(os.Stdout, &r)
	// Output:
	// You cracked the code!
}
