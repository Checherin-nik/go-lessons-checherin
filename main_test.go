package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestTrueUser(t *testing.T) {
	assert := assert.New(t)

	assert.True(trueUser("Николай", []string{"Николай", "Николас", "Николя"}), 
	"Имя 'Николай' есть в списке, ожидаем успех")

	assert.False(trueUser("Шайтан", []string{"Дьявол", "Сатана", "Чертила"}),
	"Имени 'Шайтан' нет в списке, ожидаем ошибку")
	
	assert.False(trueUser("Игорь", []string{}), 
	"Список пустой, функция должна вернуть false")

	assert.False(trueUser(" ", []string{"Ева", "Ада", "Оля"}),
	"Пробел должен возвращать ошибку")
}
