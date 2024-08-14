package main

import "fmt"

func trueUser(name string, validNames [5]string) {
	for _, validName := range validNames {
		if name == validName {
			fmt.Println("Рады вас приветствовать, " + name + "!")
			return
		}
	}
	fmt.Println("Пошёл на хуй, " + name)
}

func main() {

	var validNames [5]string
	validNames[0] = "Николай"
	validNames[1] = "Егор"
	validNames[2] = "Иван"
	validNames[3] = "Андрей"
	validNames[4] = "Тормунд"

	var greeting string
	fmt.Print("Введите ваше имя: ") 
	fmt.Scan(&greeting)

	trueUser(greeting, validNames)
}
