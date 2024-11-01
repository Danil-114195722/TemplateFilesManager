package run_funcs

import (
	"fmt"
	"os"
	"os/exec"
	
	"github.com/spf13/cobra"

	"github.com/Danil-114195722/TemplateFilesManager/db"
	"github.com/Danil-114195722/TemplateFilesManager/db/models"
	"github.com/Danil-114195722/TemplateFilesManager/settings"
)

func EditRunE(cmd *cobra.Command, args []string) error {
	// парсинг флагов имени и тега нового файла
	nameFlagValue, err := cmd.Flags().GetString("name")
	if err != nil {
		return err
	}
	tagFlagValue, err := cmd.Flags().GetString("tag")
	if err != nil {
		return err
	}

	// если такого файла с таким тегом нет в БД
	var file models.File
	dbConnect := db.GetConnection()
	selectResult := dbConnect.Where("name = ? AND tag = ?", nameFlagValue, tagFlagValue).First(&file)
	if selectResult.Error != nil {
		fmt.Println("___Warning___")
		fmt.Println("File-template with such name and tag does not exist!")
		fmt.Printf("\nUse «template add -n %s -t %s» to create new file-template\n", nameFlagValue, tagFlagValue)
		return nil
	}

	// получение текстового редактора по умолчанию
	selectEditor, _ := settings.GetSelectedEditor()

	// создание команды для запуска редактора файла-шаблона
	command := exec.Command(selectEditor, fmt.Sprintf("%s/%s/%s", settings.FilesPath, nameFlagValue, tagFlagValue))
	// для корректной передачи управления из утилиты в редактор файла-шаблона
	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	// выполнение команды
	err = command.Run()
	if err != nil {
		return err
	}

	return nil
}
