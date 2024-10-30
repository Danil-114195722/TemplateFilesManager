package cmd

import (
	"os"

	"github.com/spf13/cobra"
)


// запуск утилиты без команд, аргументов и флагов
var rootCmd = &cobra.Command{
	Use:	"template",
	Short:	"Template files manager for CLI",
	Long:	`Template files manager for CLI.
  This utility allows you to create, edit and paste your template files somewhere you want.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.TemplateFilesManager.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


