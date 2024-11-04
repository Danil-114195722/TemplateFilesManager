package cmd

import (
	"github.com/spf13/cobra"

	"github.com/ej-you/TemplateFilesManager/run_funcs"
)


// копирование файла-шаблона в место назначения
var cpCmd = &cobra.Command{
	Use:   "cp",
	Short: "Copy file-template to destination (current path)",
	RunE: run_funcs.CpRunE,
}

func init() {
	rootCmd.AddCommand(cpCmd)

	// флаг для определения имени файла
	cpCmd.Flags().StringP("name", "n", "", "Name for new file-template (required)")
	cpCmd.MarkFlagRequired("name")
	// флаг для определения тега файла, по умолчанию - default
	cpCmd.Flags().StringP("tag", "t", "default", "Tag for new file-template (optional)")
}
