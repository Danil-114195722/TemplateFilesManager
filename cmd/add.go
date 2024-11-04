package cmd

import (
	"github.com/spf13/cobra"

	"github.com/ej-you/TemplateFilesManager/run_funcs"
)

// добавление нового файла-шаблона
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add new file-template",
	RunE: run_funcs.AddRunE,
}

func init() {
	rootCmd.AddCommand(addCmd)

	// флаг для определения имени файла
	addCmd.Flags().StringP("name", "n", "", "Name for new file-template (required)")
	addCmd.MarkFlagRequired("name")
	// флаг для определения тега файла, по умолчанию - default
	addCmd.Flags().StringP("tag", "t", "default", "Tag for new file-template (optional)")
}
