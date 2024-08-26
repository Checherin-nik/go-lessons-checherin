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

func TestDeleteName(t *testing.T) {
	assert := assert.New(t)

	names := []string{"Адам", "Бен", "Веном", "Гавриил", "Дебошир"}
	
	modifiedNames := deleteName(append([]string(nil), names...), "Адам")
	expected := []string{"Бен", "Веном", "Гавриил", "Дебошир"}
	assert.Equal(expected, modifiedNames, "Имя 'Адам' есть в списке, ожидаем удаление")

	modifiedNames = deleteName(append([]string(nil), names...), "Михаил")
	assert.Equal(names, modifiedNames, "Имени 'Михаил' нет в списке, ожидаем, что ничего не изменится")

	modifiedNames = deleteName(append([]string(nil), names...), "Веном")
	expected = []string{"Адам", "Бен", "Гавриил", "Дебошир"}
	assert.Equal(expected, modifiedNames, "Удаляем имя из середины списка, ожидаем сохранение сортировки")

	modifiedNames = deleteName([]string{"Андрей"}, "Андрей")
	assert.Empty(modifiedNames, "Ожидаем пустой список после удаления последнего имени")
}


func TestAddName(t *testing.T) {
	assert := assert.New(t)

	names := []string{"Адам", "Бен", "Веном", "Гавриил", "Дебошир"}

	modifiedNames := addName(names, "Андрей")
	expected := []string{"Адам", "Андрей", "Бен", "Веном", "Гавриил", "Дебошир"}
	assert.Equal(expected, modifiedNames, "Имени 'Андрей' нет в списке, ожидаем добавления имени в список")

	modifiedNames = addName(names, "Дебошир")
	assert.Equal(names, modifiedNames, "Имени 'Дебошир' нет в списке, ожидаем что список не должен измениться")

	modifiedNames = addName([]string{}, "Андрей")
	expected = []string{"Андрей"}
	assert.Equal(expected, modifiedNames, "Ожидаем список из одного имени после добавления в пусто список")

	modifiedNames = addName([]string{"Адам"}, "Андрей")
	expected = []string{"Адам", "Андрей"}
	assert.Equal(expected, modifiedNames, "Ожидаем отсортированный список после добавления элемента")
}

func TestInitialize(t *testing.T) {
	assert := assert.New(t)

	names := []string{"Артем", "Борат", "Вениамин", "Геннадий", "Дерьмодемон"}
	expected := []string{"Артем", "Борат", "Вениамин", "Геннадий", "Дерьмодемон"}
	assert.Equal(expected, initialize(names), "Ожидаем тот же, уже отсортированный, список")

	names = []string{"Дерьмодемон", "Геннадий", "Вениамин", "Борат", "Артем"}
	expected = []string{"Артем", "Борат", "Вениамин", "Геннадий", "Дерьмодемон"}
	assert.Equal(expected, initialize(names), "Ожидаем отсоритрованный список")

	names = []string{"Damodred", "Ieremia", "Velzevul", "Gakadarion", "Abaddon"}
	expected = []string{"Abaddon", "Damodred", "Gakadarion", "Ieremia", "Velzevul"}
	assert.Equal(expected, initialize(names), "Ожидаем, что латиница тоже сортируется")
}