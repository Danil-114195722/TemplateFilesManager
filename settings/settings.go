package settings

import (
	"errors"
	"fmt"
	"os"
	"regexp"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)


// загрузка переменных окружения
var _ error = godotenv.Load("")


var homePath string = os.Getenv("HOME")
var BaseDir string = homePath + "/.local/share/TemplateFilesManager"

// путь до файлов установленной утилиты
var FilesPath string = BaseDir + "/files"

// текущая директория, откуда запущена утилита
var CurrentPath string = func() string {
	path, err := os.Getwd()
	if err != nil {
		return ""
	}
	return path
}()


// функция печати ошибки (красный цвет сообщения)
var ErrorPrintf = color.New(color.FgRed).PrintfFunc()
// функция печати предупреждения и подсказки (яркий жёлтый цвет сообщения)
var WarningPrintf = color.New(color.FgHiYellow).PrintfFunc()
// функция печати успеха (зелёный цвет сообщения)
var SuccessPrintf = color.New(color.FgGreen).PrintfFunc()
// функция печати подсказки (жёлтый цвет сообщения)
var HintPrintf = color.New(color.FgYellow).PrintfFunc()


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

// проверка наличия текущего пути, откуда запускается утилита
func CurrentPathCheck() error {
	if CurrentPath == "" {
		return errors.New("Current path was not defined!")
	}
	return nil
}
