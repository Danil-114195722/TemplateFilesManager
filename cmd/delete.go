package cmd

import (
	"github.com/spf13/cobra"

	"github.com/ej-you/TemplateFilesManager/run_funcs"
)

// удаление файлов-шаблонов и их тегов
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete files-templates and its tags",
	Long: `Delete files-templates and its tags

  You must use at least --all or --name flag. The --tag flag is only available with the --name flag.

  Using «template delete -a» will delete all files-templates.
  Using «template delete -n <filename> -a» will delete all tags of file-template with name <filename> (full delete of file-template)
  
  Using «template delete -n <filename>» will delete default tag of file-template with name <filename>
  Using «template delete -n <filename> -t <tag>» will delete tag named <tag> of file-template with name <filename>`,
	RunE: run_funcs.DeleteRunE,
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// флаг для определения удаления всех файлов-шаблонов (всех тегов одного файла шаблона)
	deleteCmd.Flags().BoolP("all", "a", false, "Delete all entities (optional)")
	// флаг для определения имени файла
	deleteCmd.Flags().StringP("name", "n", "", "File-template name (optional)")
	// флаг для определения тега файла, по умолчанию - default
	deleteCmd.Flags().StringP("tag", "t", "default", "File-template tag (optional)")

	// ставим обязательным использование одного из флагов all или name
	deleteCmd.MarkFlagsOneRequired("all", "name")
}
