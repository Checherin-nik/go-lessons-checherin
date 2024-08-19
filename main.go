package main

import ("fmt"
		"strings")

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

		for {
			fmt.Print("Введите команду: ")
			var command string
			fmt.Scan(&command)

			if command == "exit" {
				fmt.Println("Вы закончили, " + greeting)
				break
			} else if command == "print" {
				fmt.Println("Список доступных имён: " + strings.Join(validNames, ", "))
			} else if command == "add" {
				fmt.Print("Введите имя, которое хотите добавить: ")
				var newName string
				fmt.Scan(&newName)

				if trueUser(newName, validNames) {
					fmt.Println("Имя уже существует в списке.")
				} else {
					validNames = append(validNames, newName)
					fmt.Println("Имя " + newName + " добавлено в список.")
				}
			}
		}
	} else {
		fmt.Println("Пошёл на хуй, " + greeting)
	}
}