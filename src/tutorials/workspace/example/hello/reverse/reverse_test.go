// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package reverse

import "testing"

func TestString_DeveReverterConjuntoDados(t *testing.T) {
	// given
	datasForTest := []struct{ input, output string }{
		{"Hello, world", "dlrow ,olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	}

	for _, data := range datasForTest {
		// when
		got := String(data.input)

		// then
		if got != data.output {
			t.Errorf("String(%q) == %q, want %q", data.input, got, data.output)
		}
	}
}

func TestString_DeveReverterString(t *testing.T) {
	// given
	var input, saidaEsperada string = "hello", "olleh"

	// when
	stringReversed := String(input)

	// then - FAIL, if condition TRUE.
	if stringReversed != saidaEsperada {
		t.Errorf("reverse.String(%q) == %q, esperava %q", input, stringReversed, saidaEsperada)
	}
}
