package main

import (
	"fmt"
	"strings"
	"checherin-lessons-go/users"
	"checherin-lessons-go/listUsers"
)

const filename = "list.txt"

func main() {
	var usersList []string
	var err error

	usersList, err = listUsers.Load(filename)
	if err != nil {
		fmt.Println("Ошибка при загрузке списка", err)
		usersList = []string{"Николай", "Егор", "Иван", "Андрей", "Тормунд"}
	}

	names := &users.Names{}
	names.Initialize(usersList)

reLogin:

	for {
		var greeting string
		fmt.Print("Введите ваше имя: ")
		fmt.Scan(&greeting)

		if !names.TrueUser(greeting) {
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

				err := names.Add(newUser)
				if err != nil {
					fmt.Println("Ошибка при добавлении имени", err)
				} else {
					fmt.Println("Имя " + newUser + " добавлено в список.")
				}

			case "delete":
				fmt.Print("Введите имя, которое хотите удалить: ")
				var userToDelete string
				fmt.Scan(&userToDelete)

				err := names.Delete(userToDelete)
				if err != nil {
					fmt.Println("Ошибка при удалении имени:", err)
				} else {
					fmt.Println("Имя " + userToDelete + " удалено из списка.")
				}

			case "save":
				err := listUsers.Save(filename, usersList)
				if err != nil {
					fmt.Println("Ошибка при создании файла:", err)
				} else {
					fmt.Println("Список успешно сохранен в файл " + filename)
				}
				
			default:
				fmt.Println("Неизвестная команда: " + command)
			}
		}
	}
}