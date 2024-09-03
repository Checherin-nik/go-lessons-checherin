package listUsers

import (
	"os"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	assert := assert.New(t)

	t.Run("Успешная загрузка файла", func(t *testing.T) {
		filename := "testfile.txt"
		content := "Николай\nМихаил\nГавриил\nУриил"
		err := os.WriteFile(filename, []byte(content), 0644)
		assert.NoError(err, "Файл успешно создается")

		defer os.Remove(filename)

		users, err := Load(filename)
		expected := []string{"Николай", "Михаил", "Гавриил", "Уриил"}
		assert.NoError(err, "Загрузка файла прошла без ошибок")
		assert.Equal(expected, users, "Содержимое файла соответствует ожидаемому")
	})

	t.Run("Файл не найден", func(t *testing.T) {
		users, err := Load("nonexistent.txt")
		assert.Nil(users, "Список должен быть nil при ошибке")
		assert.Error(err, "Должна возвращаться ошибка при отсутствии файла")
	})

	t.Run("Пустой файл", func(t *testing.T) {
		filename := "emptyfile.txt"
		err := os.WriteFile(filename, []byte(""), 0644)
		assert.NoError(err, "Пустой файл должен успешно создаться")

		defer os.Remove(filename)

		users, err := Load(filename)
		assert.NoError(err, "Загрузка пустого файла должна пройти без ошибок")
		assert.Empty(users, "Список должен быть пустым для пустого файла")
	})
}