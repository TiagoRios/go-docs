package methodsinterfaces

func ExampleMinhaInterface_do_int() {
	var i interface{} = 14
	do(i)
	// Output: Twice 14 is 28
}

func ExampleMinhaInterface_do_string() {
	var i interface{} = "hello"
	do(i)
	// Output: "hello" is 5 bytes long
}

func ExampleMinhaInterface_do_unknown() {
	var i interface{} = []string{`a`, `b`}
	do(i)
	// Output: I don't know about type []string!
}
