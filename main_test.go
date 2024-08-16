package main

import (
	"testing"
)

func TestTrueUser(t *testing.T) {
	if !trueUser("Николай", []string{"Николай", "Николас", "Николя"}) {
		t.Errorf("Имя 'Николай' есть в списке, ожидаем успех")
	}

	if trueUser("Шайтан", []string{"Дьявол", "Сатана", "Чертила"}) {
		t.Errorf("Имени 'Шайтан' нет в списке, ожидаем ошибку")
	}

	if trueUser("Игорь", []string{}) {
		t.Errorf("Список пустой, функция должна вернуть false")
	}	

	if trueUser(" ", []string{"Ева", "Ада", "Оля"}) {
		t.Errorf("Пробел должен возвращать ошибку")
	}
}
