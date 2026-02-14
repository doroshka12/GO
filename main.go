package main

import "fmt"

func main() {
    fmt.Println("Привет, мир!")
    fmt.Println("Мое первое приложение на Go!")
    
    var name string
    fmt.Print("Как тебя зовут? ")
    fmt.Scanln(&name)
    fmt.Printf("Приятно познакомиться, %s!\n", name)
}
