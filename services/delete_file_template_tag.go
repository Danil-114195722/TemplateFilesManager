package services

import (
	"errors"
	"fmt"
	"os"

	"github.com/Danil-114195722/TemplateFilesManager/settings"
)


// удаление файла тега и директории файла-шаблона, если тегов не осталось
func DeleteFileTemplateTag(name, tag string) error {
	// удаление файла тега
	err := os.Remove(fmt.Sprintf("%s/%s/%s", settings.FilesPath, name, tag))
	if err != nil {
		return err
	}

	// получаем все теги файла-шаблона
	fileTemplateTags, err := FindFileTemplateTags(name, "")
	if err != nil {
		return err
	}
	// если теперь у файла-шаблона не осталось тегов, то удаляем его директорию
	if len(fileTemplateTags) == 0 {
		err = os.Remove(fmt.Sprintf("%s/%s", settings.FilesPath, name))
		// проверка на ошибку удаления НЕпустой директории
		if err != nil {
			return err
		}
	}
	return nil
}

// полное удаление директории файла-шаблона со всеми тегами
func DeleteFileTemplate(name string) error {
	// удаление директории файла-шаблона
	err := os.RemoveAll(fmt.Sprintf("%s/%s", settings.FilesPath, name))
	if err != nil {
		return err
	}
	return nil
}

// удаление всех директорий файлов-шаблонов
func DeleteAllFilesTemplates() (int, error) {
	// получение всех имён директорий файлов-шаблонов
	allFilesTemplates, err := FindFilesTemplates("")
	if err != nil {
		return 0, err
	}

	// если файлов-шаблонов не найдено
	if len(allFilesTemplates) == 0 {
		return 0, errors.New("files-templates not found")
	}

	// кол-во успешно удалённых файлов-шаблонов
	successDeletedCount := 0
	var removeDirErr error
	for _, fileTemplate := range allFilesTemplates {
		// удаление директории файла-шаблона
		err := os.RemoveAll(fmt.Sprintf("%s/%s", settings.FilesPath, fileTemplate))
		// если происходит ошибка при удалении какого-то файла-шаблона, то продолжаем удалять остальное
		if err != nil {
			removeDirErr = err
		} else {
			successDeletedCount++
		}
	}
	if removeDirErr != nil {
		return successDeletedCount, removeDirErr
	}
	return successDeletedCount, nil
}
