package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/Danil-114195722/TemplateFilesManager/settings"
)


// запуск утилиты без команд, аргументов и флагов
var rootCmd = &cobra.Command{
	Use:	"template",
	Short:	"Files-templates manager for CLI",
	Long:	`Files-templates manager for CLI.
  This utility allows you to create, edit, delete and cp your template files somewhere you want.

  Created files are stored in "~/.local/share/TemplateFilesManager/files" directory.
  Each stored file-template is a directory containing a list of files named by file-template tag names.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		settings.ErrorPrintf("Error: %s\n\n", err.Error())

		// поиск подкоманды, в которой произошла ошибка, по предоставленным утилите аргументам
        currentCommand, _, err := rootCmd.Find(os.Args[1:])
        if err == nil { // NOT err
			// выводим справку для подкоманды, в которой произошла ошибка
			currentCommand.Usage()
        }
		os.Exit(1)
	}
}

func init() {
	// отключение встроенного вывода ошибки
	rootCmd.SilenceErrors = true
	// отключение встроенного вывода справки об использовании при ошибке
	rootCmd.SilenceUsage = true
}