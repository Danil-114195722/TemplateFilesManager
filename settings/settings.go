package settings

import (
	"errors"
	"fmt"
	"os"
	"regexp"

	"github.com/joho/godotenv"
)


// загрузка переменных окружения
var _ error = godotenv.Load("./settings/config/.env")


var homePath string = os.Getenv("HOME")
var BaseDir string = homePath + "/.local/share/TemplateFilesManager"

// путь до файлов установленной утилиты
var FilesPath string = BaseDir + "/files"


// получение выбранного редактора по умолчанию
func GetSelectedEditor() (string, error) {
	var emptyString string

	// чтение файла, содержащего выбранный редактор по умолчанию
	filePath := homePath + "/.selected_editor"
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		return emptyString, errors.New(fmt.Sprintf("File %q was not found!", filePath))
	}
	// парсинг пути к выбранному редактору из файла
	// при найденной подстроке будет такое содержание reFoundSlice = [SELECTED_EDITOR="/bin/nano" /bin/nano]
	reFoundSlice := regexp.MustCompile(`(?m)^SELECTED_EDITOR="?(.+?)"?$`).FindStringSubmatch(string(fileData))

	if len(reFoundSlice) != 2 {
		return emptyString, errors.New(fmt.Sprintf("Selected default editor was not found in file %q!", filePath))
	}
	return reFoundSlice[1], nil
}

// проверка наличия переменной окружения HOME
func HomeEnvCheck() error {
	if homePath == "" {
		return errors.New("HOME env was not found!")
	}
	return nil
}
