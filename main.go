package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
	"checherin-lessons-go/users"
)

func main() {
	var usersList []string

	if _, err := os.Stat("list.txt"); err == nil {
		file, _ := os.Open("list.txt")
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			usersList = append(usersList, scanner.Text())
		}

		fmt.Println("Список загружен из файла list.txt")
	} else {
		usersList = []string{"Николай", "Егор", "Иван", "Андрей", "Тормунд"}
	}

	usersList = users.Initialize(usersList)

reLogin:

	for {
		var greeting string
		fmt.Print("Введите ваше имя: ")
		fmt.Scan(&greeting)

		if !users.TrueUser(greeting, usersList) {
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

			case "logout":
				fmt.Println("Вы вышли из системы, " + greeting)
				goto reLogin

			case "print":
				fmt.Println("Список доступных имён: " + strings.Join(usersList, ", "))

			case "add":
				fmt.Print("Введите имя, которое хотите добавить: ")
				var newUser string
				fmt.Scan(&newUser)

				if users.TrueUser(newUser, usersList) {
					fmt.Println("Имя уже существует в списке.")
				} else {
					usersList = users.AddName(usersList, newUser)
					fmt.Println("Имя " + newUser + " добавлено в список.")
				}

			case "delete":
				fmt.Print("Введите имя, которое хотите удалить: ")
				var userToDelete string
				fmt.Scan(&userToDelete)

				if users.TrueUser(userToDelete, usersList) {
					usersList = users.DeleteName(usersList, userToDelete) 
					fmt.Println("Имя " + userToDelete + " удалено из списка")
				} else {
					fmt.Println("Имя не найдено в списке")
				}

			case "save":
				file, _ := os.Create("list.txt")
				defer file.Close()

				for _, user := range usersList {
					file.WriteString(user + "\n")
					}
				fmt.Println("Список успешно сохранен в файл list.txt")

			default:
				fmt.Println("Неизвестная команда: " + command)
			}
		}
	}
}