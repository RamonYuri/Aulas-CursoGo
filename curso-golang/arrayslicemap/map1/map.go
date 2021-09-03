package main

import "fmt"

func main() {
	// var aprovados map[int]string
	// maps devem ser inicializados
	aprovados := make(map[int]string)

	aprovados[12345678978] = "Maria"
	aprovados[98765432100] = "Pedro"
	aprovados[95175345682] = "Ana"

	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d\n", nome, cpf)
	}

	fmt.Println(aprovados[95175345682])
	delete(aprovados, 95175345682)
	fmt.Println(aprovados[95175345682])
}