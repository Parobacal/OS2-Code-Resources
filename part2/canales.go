package main

import (
	"fmt"
	//"time"
)

/* descomentar para segunda parte

func main() {
    c := make(chan string, 50)
    go count("sheep", c)

    fmt.Println("Esperando que el canal se llene")
    for msg := range c {
        fmt.Println(msg)
    }
}


func count(thing string, c chan string) {
    for i := 1; i <= 50; i++ {
        fmt.Print(".")
        c <- thing
        time.Sleep(time.Millisecond * 500)
    }
    close(c)
}
*/

/*primera parte*/

func main() {
	c := make(chan string, 1) //declaracion
	c <- "algo"               //asignacion de valor
	fmt.Println(<-c)          //esperar valor
}
