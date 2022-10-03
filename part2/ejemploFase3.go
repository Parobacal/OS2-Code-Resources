package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

func trabajador(parte string) {
	log.Println(parte, "es la parte asignada al trabajador n...n-1 ")
	time.Sleep(time.Duration(rand.Int63n(1e6)))
	log.Println("la parte", parte, "del trabajador n...n-1 completa el trabajo")
	wg.Done()
}

var (
	arregloPartes = []string{"A", "B", "C", "D"}
	ciclos        = 3
	wg            sync.WaitGroup
)

func main() {
	rand.Seed(time.Now().UnixNano())
	for c := 1; c <= ciclos; c++ {
		log.Println("inicia el ciclo de trabajo", c)
		wg.Add(len(arregloPartes))
		for _, parte := range arregloPartes {
			go trabajador(parte) // < ==== Aquí implementamos el "FORK"
		}
		wg.Wait() // < ==== Aquí implementamos el "JOIN"
		log.Println("Ciclo de tabajo ", c, " completo")
	}
}
