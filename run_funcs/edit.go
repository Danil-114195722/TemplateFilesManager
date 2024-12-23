package run_funcs

import (
	"fmt"
	"os"
	"os/exec"
	
	"github.com/spf13/cobra"

	"github.com/ej-you/TemplateFilesManager/services"
	"github.com/ej-you/TemplateFilesManager/settings"
)

func EditRunE(cmd *cobra.Command, args []string) error {
	// парсинг флагов имени и тега файла
	nameFlagValue, err := cmd.Flags().GetString("name")
	if err != nil {
		return err
	}
	tagFlagValue, err := cmd.Flags().GetString("tag")
	if err != nil {
		return err
	}

	isExists, err := services.FileTemplateTagIsExists(nameFlagValue, tagFlagValue)
	if err != nil {
		return err
	}
	// если такого файла с таким тегом нет в файловой системе
	if !isExists {
		settings.WarningPrintf("WARNING: file-template with such name and tag does not exist!\n")
		settings.HintPrintf("\nHINT: use «template add -n %s -t %s» to create new file-template\n", nameFlagValue, tagFlagValue)
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
