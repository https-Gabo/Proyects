package main

import (
	"errors"
	"fmt"
)

func DividirEntero(dividendo, divisor int) (cociente, resto int, err error) {
	if divisor == 0 {
		err = errors.New("divisor no puede ser cero")
		return
	}
	cociente = dividendo / divisor
	resto = dividendo % divisor
	return
}

func main() {
	cociente, resto, err := DividirEntero(10, 3)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Cociente: %d, Resto: %d\n", cociente, resto)
	}

	cociente, resto, err = DividirEntero(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Cociente: %d, Resto: %d\n", cociente, resto)
	}

	cociente, resto, err = DividirEntero(20, 4)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Cociente: %d, Resto: %d\n", cociente, resto)
	}
}
