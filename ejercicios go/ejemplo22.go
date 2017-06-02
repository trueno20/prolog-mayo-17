package main

import (
    "fmt"
)

func Sqrt(x float64) float64 {
    
    for i:= 1.0; i<= 10.0 ; i++{
        i := i- ((i*i)-x)/(2*i)
        fmt.Println(i)
    }
    return 0
}

func main() {
    fmt.Println(Sqrt(2))
}