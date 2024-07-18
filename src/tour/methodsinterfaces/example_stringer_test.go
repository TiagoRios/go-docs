package methodsinterfaces

import "fmt"

func ExamplePerson() {
	a := Person{"Arthur Dent", 42}
	fmt.Println(a)
	// Output:
	// Arthur Dent (42 years)
}

func ExampleIPAddr_exercicio() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
	// Output:
	// loopback: 127.0.0.1
	// googleDNS: 8.8.8.8
}
