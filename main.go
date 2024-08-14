package main

import "fmt"

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

	for i := 0; i < len(validNames); i++ {
		if greeting == validNames[i] {
			fmt.Println("Рады вас приветствовать, " + greeting + "!")
			return
		}
	}
	fmt.Println("Пошёл на хуй, " + greeting)
}
