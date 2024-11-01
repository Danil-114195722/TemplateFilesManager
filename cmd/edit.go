package cmd

import (
	"github.com/spf13/cobra"

	"github.com/Danil-114195722/TemplateFilesManager/run_funcs"
)

// изменение уже созданного файла-шаблона в текстовом редакторе
var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit already created file-template in text editor",
	RunE: run_funcs.EditRunE,
}

func init() {
	rootCmd.AddCommand(editCmd)

	// флаг для определения имени файла
	editCmd.Flags().StringP("name", "n", "", "File-template name (required)")
	editCmd.MarkFlagRequired("name")
	// флаг для определения тега файла, по умолчанию - default
	editCmd.Flags().StringP("tag", "t", "default", "File-template tag")
}
