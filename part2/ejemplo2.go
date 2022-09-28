package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	now := time.Now()
	var wg sync.WaitGroup // Declarando nuestro wait group
	wg.Add(1)             // Indicamos la cantidad de rutinas a esperar
	/*
		En lugar de llamar a go work utilizamos una funcion anonima
	*/
	go func() {
		defer wg.Done()
		work()
	}() // ---------> FORK
	wg.Wait() // JOIN <--------
	fmt.Println("Ha transcurrido ", time.Since(now))
	fmt.Println("La rutina main esperó y ahora terminó")
}

func work() {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Rutina work esta trabajando")
}
