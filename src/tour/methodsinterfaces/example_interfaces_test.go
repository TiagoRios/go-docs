package methodsinterfaces

func ExampleMeuTipo_retornandoValorETipo() {
	var i MinhaInterface = &MeuTipo{"Hello"}
	describe(i)
	// Output: (&{Hello}, *methodsinterfaces.MeuTipo)
}

func ExampleMeuTipo_retornandoValorETipo_variavelNaoIniciada() {
	var meuTipo *MeuTipo
	var i MinhaInterface = meuTipo
	describe(i)
	// Output: (<nil>, *methodsinterfaces.MeuTipo)
}

func ExampleMeuTipoFloat_retornandoValorETipo() {
	var i MinhaInterface = MeuTipoFloat(12.34)
	describe(i)
	// Output: (12.34, methodsinterfaces.MeuTipoFloat)
}

func ExampleMinhaInterface_retornandoValorETipo_nilInterface() {
	var i MinhaInterface
	describe(i)
	// Output: (<nil>, <nil>)
}

func ExampleMinhaInterface_retornandoValorETipo_emptyInterface() {
	var i interface{}
	describe(i)
	// Output: (<nil>, <nil>)
}

func ExampleMinhaInterface_retornandoValorETipo_inteiro() {
	var i interface{} = 412
	describe(i)
	// Output: (412, int)
}

func ExampleMinhaInterface_retornandoValorETipo_bool() {
	var i interface{} = false
	describe(i)
	// Output: (false, bool)
}

func ExampleMinhaInterface_retornandoValorETipo_string() {
	var i interface{} = "testando"
	describe(i)
	// Output: (testando, string)
}
