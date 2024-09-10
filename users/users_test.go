package users

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestTrueUser(t *testing.T) {
	assert := assert.New(t)

	names := Names{validNames: []string{"Николай", "Николас", "Николя"}}

	assert.True(names.TrueUser("Николай"), "Имя 'Николай' есть в списке, ожидаем успех")

	names = Names{validNames: []string{"Дьявол", "Сатана", "Чертила"}}
	assert.False(names.TrueUser("Шайтан"), "Имени 'Шайтан' нет в списке, ожидаем ошибку")
	
	names = Names{validNames: []string{}}
	assert.False(names.TrueUser("Игорь"), "Список пустой, функция должна вернуть false")

	names = Names{validNames: []string{"Ева", "Ада", "Оля"}}
	assert.False(names.TrueUser(" "),"Пробел должен возвращать ошибку")
}

func TestDeleteName(t *testing.T) {
	assert := assert.New(t)

	names := Names{validNames:[]string{"Адам", "Бен", "Веном", "Гавриил", "Дебошир"}}
	
	err := names.Delete("Адам")
	expected := []string{"Бен", "Веном", "Гавриил", "Дебошир"}
	assert.NoError(err, "Ожидаем, что ошибки не будет при удалении существующего имени")
	assert.Equal(expected, names.GetValidNames(), "Имя 'Адам' есть в списке, ожидаем удаление")

	err = names.Delete("Михаил")
	assert.ErrorIs(err, ErrNotFound, "Ожидаем ошибку ErrNotFound при удалении несуществующего имени")
	assert.Equal(expected, names.GetValidNames(), "Имени 'Михаил' нет в списке, ожидаем, что ничего не изменится")

	err = names.Delete("Веном")
	expected = []string{"Бен", "Гавриил", "Дебошир"}
	assert.NoError(err, "Ожидаем, что ошибки не будет при удалении существующего имени")
	assert.Equal(expected, names.GetValidNames(), "Удаляем имя из середины списка, ожидаем сохранение сортировки")

	names = Names{validNames: []string{"Андрей"}}
	err = names.Delete("Андрей")
	assert.NoError(err, "Ожидаем, что ошибки не будет при удалении существующего имени")
	assert.Empty(names.GetValidNames(), "Ожидаем пустой список после удаления последнего имени")
}


func TestAddName(t *testing.T) {
	assert := assert.New(t)

	names := Names{validNames: []string{"Адам", "Бен", "Веном", "Гавриил", "Дебошир"}}

	err := names.Add("Андрей")
	expected := []string{"Адам", "Андрей", "Бен", "Веном", "Гавриил", "Дебошир"}
	assert.NoError(err, "Ожидаем, что ошибки не будет при добавлении нового имени")
	assert.Equal(expected, names.GetValidNames(), "Имени 'Андрей' нет в списке, ожидаем добавления имени в список")

	err = names.Add("Дебошир")
	assert.ErrorIs(err, ErrAlreadyExist, "Ожидаем ошибку ErrAlreadyExist при удалении несуществующего имени")
	assert.Equal(expected, names.GetValidNames(), "Имени 'Дебошир' нет в списке, ожидаем что список не должен измениться")

	names = Names{validNames: []string{}}
	err = names.Add("Андрей")
	expected = []string{"Андрей"}
	assert.NoError(err, "Ожидаем, что ошибки не будет при добавлении в пустой список")
	assert.Equal(expected, names.GetValidNames(), "Ожидаем список из одного имени после добавления в пусто список")

	names = Names{validNames: []string{"Адам"}}
	err = names.Add("Андрей")
	expected = []string{"Адам", "Андрей"}
	assert.NoError(err, "Ожидаем, что ошибки не будет при добавлении нового имени")
	assert.Equal(expected, names.GetValidNames(), "Ожидаем отсортированный список после добавления элемента")
}

func TestInitialize(t *testing.T) {
	assert := assert.New(t)

	names := Names{}
	names.Initialize([]string{"Артем", "Борат", "Вениамин", "Геннадий", "Дерьмодемон"})
	expected := []string{"Артем", "Борат", "Вениамин", "Геннадий", "Дерьмодемон"}
	assert.Equal(expected, names.GetValidNames(), "Ожидаем тот же, уже отсортированный, список")

	names.Initialize([]string{"Дерьмодемон", "Геннадий", "Вениамин", "Борат", "Артем"})
	expected = []string{"Артем", "Борат", "Вениамин", "Геннадий", "Дерьмодемон"}
	assert.Equal(expected, names.GetValidNames(), "Ожидаем отсоритрованный список")

	names.Initialize([]string{"Damodred", "Ieremia", "Velzevul", "Gakadarion", "Abaddon"})
	expected = []string{"Abaddon", "Damodred", "Gakadarion", "Ieremia", "Velzevul"}
	assert.Equal(expected, names.GetValidNames(), "Ожидаем, что латиница тоже сортируется")
}