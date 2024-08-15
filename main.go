package main

import "fmt"

func trueUser(name string, validNames []string) bool {
	for _, validName := range validNames {
		if name == validName {
			return true
		}
	}
	return false
}

func main() {

	validNames := []string{"Николай", "Егор", "Иван", "Андрей", "Тормунд"}

	var greeting string
	fmt.Print("Введите ваше имя: ") 
	fmt.Scan(&greeting)

	if trueUser(greeting, validNames) {
		fmt.Println("Рады вас приветствовать, " + greeting + "!")
	} else {
		fmt.Println("Пошёл на хуй, " + greeting)
	}
}
