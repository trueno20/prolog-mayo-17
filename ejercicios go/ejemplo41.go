package main

import (
    "fmt"
    "time"
)

func main() {
    fmt.Println("?Cuando es Sabado?")
    today := time.Now().Weekday()
    switch time.Saturday {
    case today + 0:
        fmt.Println("Hoy.")
    case today + 1:
        fmt.Println("Manana.")
    case today + 2:
        fmt.Println("Pasado manana.")
    default:
        fmt.Println("Demasiado tarde.")
    }
}
