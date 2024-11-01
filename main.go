package main

import (
	"fmt"

	"github.com/Danil-114195722/TemplateFilesManager/cmd"
	"github.com/Danil-114195722/TemplateFilesManager/settings"
)


func main() {
	// проверки важных составляющих утилиты перед её запуском
	if err := settings.HomeEnvCheck(); err != nil {
		fmt.Println("Error:", err.Error())
		return
	}
	// проверки важных составляющих утилиты перед её запуском
	_, err := settings.GetSelectedEditor()
	if err != nil {
		fmt.Println("Error:", err.Error())
		fmt.Println("Try to set up default editor with command «select-editor»")
		return
	}

	cmd.Execute()
}
