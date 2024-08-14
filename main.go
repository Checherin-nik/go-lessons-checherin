package main

import "fmt"

func main() {
	var greeting string
	fmt.Print("Введите ваше имя: ") 
	fmt.Scan(&greeting)
	if greeting == "Николай" {        
		fmt.Println("Рады вас приветствовать, " + greeting + "!")
	} else {
		fmt.Println("Пошёл на хуй, " + greeting)
	}
}
