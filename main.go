package main

import (
	"github.com/ej-you/TemplateFilesManager/cmd"
	"github.com/ej-you/TemplateFilesManager/settings"
)


func main() {
	var err error

	// проверка нахождения значения HOME переменной окружения
	if err = settings.HomeEnvCheck(); err != nil {
		settings.ErrorPrintf("Error: %s\n", err.Error())
		return
	}
	// проверка определения текущего пути, откуда запущена утилита
	if err = settings.CurrentPathCheck(); err != nil {
		settings.ErrorPrintf("Error: %s\n", err.Error())
		return
	}
	// проверка выбранного редактора по умолчанию
	_, err = settings.GetSelectedEditor()
	if err != nil {
		settings.ErrorPrintf("Error: %s\n", err.Error())
		settings.HintPrintf("\nHINT: Try to set up default editor with command «select-editor»\n")
		return
	}

	cmd.Execute()
}
