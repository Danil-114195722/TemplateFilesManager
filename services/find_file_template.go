package services

import (
	"errors"
	"fmt"
	"os"
	"regexp"

	"github.com/Danil-114195722/TemplateFilesManager/settings"
)


// проверка тега файла-шаблона на существование
func FileTemplateTagIsExists(name, tag string) (bool, error) {
	fullFilePath := fmt.Sprintf("%s/%s/%s", settings.FilesPath, name, tag)

	fileInfo, err := os.Stat(fullFilePath)
	if err != nil {
		// неизвестная ошибка
		if !os.IsNotExist(err) {
			return false, err
		// файла не существует
		} else {
			return false, nil
		}
	}
	// если файл является директорией
	if fileInfo.IsDir() {
		return false, errors.New(fmt.Sprintf("Found dir instead file at %q", fullFilePath))
	}
	// файл существует
	return true, nil
}

// проверка файла-шаблона на существование
func FileTemplateIsExists(name string) (bool, error) {
	fullFilePath := fmt.Sprintf("%s/%s", settings.FilesPath, name)

	fileInfo, err := os.Stat(fullFilePath)
	if err != nil {
		// неизвестная ошибка
		if !os.IsNotExist(err) {
			return false, err
		// файла не существует
		} else {
			return false, nil
		}
	}
	// если файл НЕ является директорией
	if !fileInfo.IsDir() {
		return false, errors.New(fmt.Sprintf("Found file instead dir at %q", fullFilePath))
	}
	// файл существует
	return true, nil
}


// получение всех имён файлов в директории dirPath имеющих подстроку match
func getFilteredFilesNamesFromDir(dirPath, match string) ([]string, error) {
	var filesNamesList []string

	// открываем каталог со всеми файлами
	allFilesDir, err := os.Open(dirPath)
	if err != nil {
		return filesNamesList, err
	}

	// получение всех названий файлов из папки
	filesNames, err := allFilesDir.Readdirnames(0)
    if err != nil {
        return filesNamesList, err
    }

    // фильтрация по содержанию подстроки match
    if match != "" {
    	reFilesMatching := regexp.MustCompile(`.*?` + match + `.*?`)
    	// перебор всех имён файлов-шаблонов
    	for _, fileName := range filesNames {
    		if reFilesMatching.MatchString(fileName) {
    			filesNamesList = append(filesNamesList, fileName)
    		}
    	}
    	return filesNamesList, nil
    }
    // если фильтрация не нужна
	return filesNames, nil
}

// поиск всех файлов-шаблонов без их тегов (возможна фильтрация по подстроке match)
func FindFilesTemplates(match string) ([]string, error) {
	return getFilteredFilesNamesFromDir(settings.FilesPath, match)
}

// поиск всех тегов указанного файла-шаблона (возможна фильтрация по подстроке match)
func FindFileTemplateTags(fileName, match string) ([]string, error) {
	return getFilteredFilesNamesFromDir(fmt.Sprintf("%s/%s", settings.FilesPath, fileName), match)
}
