package main

import (
	"fmt"
	"strings"
	"slices"
)

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

	if !trueUser(greeting, validNames) {
		fmt.Println("Пошёл на хуй, " + greeting)
		return
	}

	fmt.Println("Рады вас приветствовать, " + greeting + "!")

	for {
		fmt.Print("Введите команду: ")
		var command string
		fmt.Scan(&command)

		switch command {
		case "exit":
			fmt.Println("Вы закончили, " + greeting)
			return 

		case "print":
			fmt.Println("Список доступных имён: " + strings.Join(validNames, ", "))

		case "add":
			fmt.Print("Введите имя, которое хотите добавить: ")
			var newName string
			fmt.Scan(&newName)

			if trueUser(newName, validNames) {
				fmt.Println("Имя уже существует в списке.")
			} else {
				validNames = append(validNames, newName)
				fmt.Println("Имя " + newName + " добавлено в список.")
			}

		case "delete":
			fmt.Print("Введите имя, которое хотите удалить: ")
			var nameToDelete string
			fmt.Scan(&nameToDelete)

			if trueUser(nameToDelete, validNames) {
				for i, validName := range validNames {
					if nameToDelete == validName {
						validNames = slices.Delete(validNames, i, i+1)
						break
					}
				}
				fmt.Println("Имя " + nameToDelete + " удалено из списка")
			} else {
				fmt.Println("Имя не найдено в списке")
			}

		default:
			fmt.Println("Неизвестная команда: " + command)
		}
	}
}