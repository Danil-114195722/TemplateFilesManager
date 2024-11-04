package run_funcs

import (
	"fmt"
	"os/exec"
	
	"github.com/spf13/cobra"

	"github.com/ej-you/TemplateFilesManager/services"
	"github.com/ej-you/TemplateFilesManager/settings"
)

func CpRunE(cmd *cobra.Command, args []string) error {
	// парсинг флагов имени и тега файла
	nameFlagValue, err := cmd.Flags().GetString("name")
	if err != nil {
		return err
	}
	tagFlagValue, err := cmd.Flags().GetString("tag")
	if err != nil {
		return err
	}

	// если такого файла-шаблона с таким тегом нет в файловой системе
	isExists, err := services.FileTemplateTagIsExists(nameFlagValue, tagFlagValue)
	if err != nil {
		return err
	}
	if !isExists {
		settings.WarningPrintf("WARNING: file-template with such name and tag does not exist!\n")
		fmt.Println("Nothing was copied")
		return nil
	}

	// определение полных путей для последующего копирования
	srcPath := fmt.Sprintf("%s/%s/%s", settings.FilesPath, nameFlagValue, tagFlagValue)
	destPath := fmt.Sprintf("%s/%s", settings.CurrentPath, nameFlagValue)

	// проверка на существование файла в текущей директории с названием файла-шаблона
	isExists, err = services.FileIsExists(destPath)
	if err != nil {
		return err
	}
	// если файл с таким же названием уже есть, то запрашиваем подтверждение на продолжение копирования
	if isExists {
		var continueCopying string
		settings.WarningPrintf("WARNING: file with name %q already exist in current directory!\n", nameFlagValue)
		fmt.Print("Continue copying file-template? [y, n] ")
		fmt.Scan(&continueCopying)

		if continueCopying != "y" {
			fmt.Println("Canceled")
			return nil
		}
		fmt.Println("Continue...")
	}
	
	// создание команды для копирования файла-шаблона в текущую директорию
	command := exec.Command("cp", srcPath, destPath)
	// выполнение команды
	err = command.Run()
	if err != nil {
		return err
	}
	
	settings.SuccessPrintf("Successfully copied file %q to current directory\n", nameFlagValue)
	return nil
}
