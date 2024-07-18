package methodsinterfaces

import (
	"fmt"
	"io"
	"strings"
)

func Leitor(texto string, quantidadeBytesLidos int) {
	leitor := strings.NewReader(texto)

	sliceByte := make([]byte, quantidadeBytesLidos) // length 8

	for {
		numeroBytesPopulados, err := leitor.Read(sliceByte)

		if err == io.EOF {
			break
		}

		imprimirSaidaConsole(sliceByte, numeroBytesPopulados, err)
	}
}

func imprimirSaidaConsole(slice []byte, numBytes int, err error) {
	fmt.Printf("bytes lidos = %v\n", numBytes)
	fmt.Printf("slice bytes lidos = %v\n", slice)
	fmt.Printf("conteudo bytes lido = %q\n", slice[:numBytes])
	fmt.Printf("err = %v\n", err)
}

// ---------- Exercicio rot13Reader ----------

type Rot13Reader struct {
	r io.Reader
}

func (rot Rot13Reader) Read(sliceBytes []byte) (numeroBytesLidos int, err error) {
	numeroBytesLidos, err = rot.r.Read(sliceBytes)

	if err == nil {
		for i := range sliceBytes {
			switch {
			case sliceBytes[i] >= 97 && sliceBytes[i] <= 109: // a - m
				sliceBytes[i] += 13
			case sliceBytes[i] > 109 && sliceBytes[i] <= 122: // n - z
				sliceBytes[i] -= 13
			case sliceBytes[i] >= 65 && sliceBytes[i] <= 77: // A - M
				sliceBytes[i] += 13
			case sliceBytes[i] > 77 && sliceBytes[i] <= 90: // N - Z
				sliceBytes[i] -= 13
			}
		}
	}

	return // naked return
}
