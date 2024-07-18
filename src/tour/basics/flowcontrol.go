package basics

import (
	"errors"
	"fmt"
	"log"
	"math"
	"runtime"
	"time"
)

func Somatorio(quantidade int) {
	sum := 0

	if quantidade == 0 {
		log.Fatal(errors.New("entrada igual a zero"))
	} else {
		for i := 0; i <= quantidade; i++ {
			sum += i
		}

		fmt.Printf("O somatorio de número(s) até o número %v é igual a: %v\n", quantidade, sum)
	}
}

func SomatorioNumeros(quantidade int) {
	inicial := 1
	sumImpares := 0
	sumPares := 0

	for inicial <= quantidade {
		if inicial%2 == 0 {
			sumPares += inicial
		} else {
			sumImpares += inicial
		}

		inicial += 1
	}

	fmt.Printf("Soma de número(s) Par(es) até o numero %v é igual a: %v\n", quantidade, sumPares)
	fmt.Printf("Soma de número(s) Ímpar(es) até o numero %v é igual a: %v\n", quantidade, sumImpares)
}

// ---------- if, else ----------

func Sqrt(x float64) string {
	if x < 0 {
		return fmt.Sprint(Sqrt(-x) + "i")
	}

	return fmt.Sprint(math.Sqrt(x))
}

func Pow(x, n, lim float64) float64 {
	// If with a short statement
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

// ---------- switch ----------

func QualSistemaOperacional() {
	fmt.Print("Go runs on ")
	// break usado automaticamente pela linguagem
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}

func QuandoESabado() {
	fmt.Println("When's Saturday?")

	today := time.Now().Weekday()

	switch time.Saturday {
	case today + 6:
		fmt.Println("Was yesterday, today is", today)
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

func Saudacao() {
	t := time.Now()

	// Switch with no condition
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

// ---------- defer ----------

func Adiar() {
	defer fmt.Println("primeiro")
	defer fmt.Println("segundo")
	defer fmt.Println("terceiro")

	fmt.Println("quarto")
	fmt.Println("ultimo")
}

func ContagemDefer() {
	fmt.Println("counting")

	for i := 0; i <= 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
