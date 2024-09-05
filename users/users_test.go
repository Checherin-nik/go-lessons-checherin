package users

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestTrueUser(t *testing.T) {
	assert := assert.New(t)

	assert.True(TrueUser("Николай", []string{"Николай", "Николас", "Николя"}), 
	"Имя 'Николай' есть в списке, ожидаем успех")

	assert.False(TrueUser("Шайтан", []string{"Дьявол", "Сатана", "Чертила"}),
	"Имени 'Шайтан' нет в списке, ожидаем ошибку")
	
	assert.False(TrueUser("Игорь", []string{}), 
	"Список пустой, функция должна вернуть false")

	assert.False(TrueUser(" ", []string{"Ева", "Ада", "Оля"}),
	"Пробел должен возвращать ошибку")
}

func TestDeleteName(t *testing.T) {
	assert := assert.New(t)

	names := []string{"Адам", "Бен", "Веном", "Гавриил", "Дебошир"}
	
	modifiedNames, err := DeleteName(append([]string(nil), names...), "Адам")
	expected := []string{"Бен", "Веном", "Гавриил", "Дебошир"}
	assert.NoError(err, "Ожидаем, что ошибки не будет при удалении существующего имени")
	assert.Equal(expected, modifiedNames, "Имя 'Адам' есть в списке, ожидаем удаление")

	modifiedNames, err = DeleteName(append([]string(nil), names...), "Михаил")
	assert.ErrorIs(err, ErrNotFound, "Ожидаем ошибку ErrNotFound при удалении несуществующего имени")
	assert.Equal(names, modifiedNames, "Имени 'Михаил' нет в списке, ожидаем, что ничего не изменится")

	modifiedNames, err = DeleteName(append([]string(nil), names...), "Веном")
	expected = []string{"Адам", "Бен", "Гавриил", "Дебошир"}
	assert.NoError(err, "Ожидаем, что ошибки не будет при удалении существующего имени")
	assert.Equal(expected, modifiedNames, "Удаляем имя из середины списка, ожидаем сохранение сортировки")

	modifiedNames, err = DeleteName([]string{"Андрей"}, "Андрей")
	assert.NoError(err, "Ожидаем, что ошибки не будет при удалении существующего имени")
	assert.Empty(modifiedNames, "Ожидаем пустой список после удаления последнего имени")
}


func TestAddName(t *testing.T) {
	assert := assert.New(t)

	names := []string{"Адам", "Бен", "Веном", "Гавриил", "Дебошир"}

	modifiedNames, err := AddName(names, "Андрей")
	expected := []string{"Адам", "Андрей", "Бен", "Веном", "Гавриил", "Дебошир"}
	assert.NoError(err, "Ожидаем, что ошибки не будет при добавлении нового имени")
	assert.Equal(expected, modifiedNames, "Имени 'Андрей' нет в списке, ожидаем добавления имени в список")

	modifiedNames, err = AddName(names, "Дебошир")
	assert.ErrorIs(err, ErrAlreadyExist, "Ожидаем ошибку ErrAlreadyExist при удалении несуществующего имени")
	assert.Equal(names, modifiedNames, "Имени 'Дебошир' нет в списке, ожидаем что список не должен измениться")

	modifiedNames, err = AddName([]string{}, "Андрей")
	expected = []string{"Андрей"}
	assert.NoError(err, "Ожидаем, что ошибки не будет при добавлении в пустой список")
	assert.Equal(expected, modifiedNames, "Ожидаем список из одного имени после добавления в пусто список")

	modifiedNames, err = AddName([]string{"Адам"}, "Андрей")
	expected = []string{"Адам", "Андрей"}
	assert.NoError(err, "Ожидаем, что ошибки не будет при добавлении нового имени")
	assert.Equal(expected, modifiedNames, "Ожидаем отсортированный список после добавления элемента")
}

func TestInitialize(t *testing.T) {
	assert := assert.New(t)

	names := []string{"Артем", "Борат", "Вениамин", "Геннадий", "Дерьмодемон"}
	expected := []string{"Артем", "Борат", "Вениамин", "Геннадий", "Дерьмодемон"}
	assert.Equal(expected, Initialize(names), "Ожидаем тот же, уже отсортированный, список")

	names = []string{"Дерьмодемон", "Геннадий", "Вениамин", "Борат", "Артем"}
	expected = []string{"Артем", "Борат", "Вениамин", "Геннадий", "Дерьмодемон"}
	assert.Equal(expected, Initialize(names), "Ожидаем отсоритрованный список")

	names = []string{"Damodred", "Ieremia", "Velzevul", "Gakadarion", "Abaddon"}
	expected = []string{"Abaddon", "Damodred", "Gakadarion", "Ieremia", "Velzevul"}
	assert.Equal(expected, Initialize(names), "Ожидаем, что латиница тоже сортируется")
}