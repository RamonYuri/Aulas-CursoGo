package main

import "fmt"

func fatorial(n uint) (uint) {
	switch {
	case n == 0:
		return 1
	default:
		return n * fatorial(n-1)
	}
}

func main() {
	resultado1 := fatorial(3)
	fmt.Println(resultado1)



}