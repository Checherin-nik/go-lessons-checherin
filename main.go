package main

import "fmt"

var greeting string

func main() {
	fmt.Print("Введите ваше имя: ") 
	fmt.Scanln(&greeting)        
	fmt.Println("Рады вас приветствовать, " + greeting + "!")
}
