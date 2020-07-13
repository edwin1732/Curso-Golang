package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	go miNombre("Edwin Valdez") //Ejecuta el nombre de manera asincrona
	fmt.Println("Estoy aqui")
	var x string
	fmt.Scanln(&x)
}

func miNombre(nombre string) {
	letras := strings.Split(nombre, "")
	for _, letra := range letras {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letra)
	}
}