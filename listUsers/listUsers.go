package listUsers

import (
	"bufio"
	"os"
	"fmt"
)

func Load(filename string) ([]string, error) {
	var usersList []string

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		usersList = append(usersList, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	fmt.Println("Список загружен из файла", filename)
	return usersList, nil
}

func Save(filename string, usersList []string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, user := range usersList {
		_, err := file.WriteString(user + "\n")
		if err != nil {
			return err
		}
	}
	return nil
}