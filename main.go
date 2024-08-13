package main

import "fmt"

func main() {
	var greeting string
	fmt.Print("Введите ваше имя: ") 
	fmt.Scanln(&greeting)        
	fmt.Println("Рады вас приветствовать, " + greeting + "!")
}
