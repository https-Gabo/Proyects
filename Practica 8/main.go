package main

import (
	"fmt"
	"sync"
)

func calcularFactorial(n uint64, wg *sync.WaitGroup, ch chan<- string) {
	defer wg.Done()

	fact := uint64(1)
	for i := uint64(2); i <= n; i++ {
		fact *= i
	}

	ch <- fmt.Sprintf("Factorial de %d: %d", n, fact)
}

func main() {
	numeros := []uint64{5, 7, 10, 15, 20}

	var wg sync.WaitGroup
	resultados := make(chan string, len(numeros))

	for _, num := range numeros {
		wg.Add(1)
		go calcularFactorial(num, &wg, resultados)
	}

	go func() {
		wg.Wait()
		close(resultados)
	}()

	fmt.Println("Resultados:")
	for res := range resultados {
		fmt.Println(res)
	}
}
