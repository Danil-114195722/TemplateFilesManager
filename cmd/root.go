package cmd

import (
	"os"

	"github.com/spf13/cobra"
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
		os.Exit(1)
	}
}
