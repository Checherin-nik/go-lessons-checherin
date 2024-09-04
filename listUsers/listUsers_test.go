package listUsers

import (
	"os"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestLoadAndSave(t *testing.T) {
	assert := assert.New(t)

	t.Run("Успешное сохранение и загрузка файла", func(t *testing.T) {
		tempFile, err := os.CreateTemp("", "testfile*.txt")
		assert.NoError(err, "Временный файл должен быть создан без ошибок")
		defer os.Remove(tempFile.Name())

		usersList := []string{"Николай", "Михаил", "Гавриил", "Уриил"}

		err = Save(tempFile.Name(), usersList)
		assert.NoError(err, "Сохранение прошло без ошибок")

		loadeUsers, err := Load(tempFile.Name())
		assert.NoError(err, "Загрузка файла прошла без ошибок")
		assert.Equal(usersList, loadeUsers, "Содержимое файла после загрузки соответствует сохраненному списку")
	})

	t.Run("Файл не найден при загрузке", func(t *testing.T) {
		loadeUsers, err := Load("nonexsistent.txt")
		assert.Nil(loadeUsers, "Список должен быть nil при ошибке")
		assert.Error(err, "Должна возвращаться ошибка при отсутствии файла")
	})

	t.Run("Загрузка и сохранение пустого файла", func(t *testing.T) {
		tempFile, err := os.CreateTemp("", "emptyfile*.txt")
		assert.NoError(err, "Временный файл должен быть создан без ошибок")
		defer os.Remove(tempFile.Name())

		emptyList := []string{}

		err = Save(tempFile.Name(), emptyList)
		assert.NoError(err, "Сохранение пустого файла должно пройти без ошибок")

		loadeUsers, err := Load(tempFile.Name())
		assert.NoError(err, "Загрузка пустого файла должна пройти без ошибок")
		assert.Empty(loadeUsers, "Список должен быть пустым для пустого файла")
	})
}