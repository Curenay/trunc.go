package main

import (
    "fmt"
)

func main() {
    var number int
    fmt.Print("Ingrese numero flotante:\n")
    fmt.Scanln(&number)
    fmt.Print("\nEntero convertido:\n",number)
}