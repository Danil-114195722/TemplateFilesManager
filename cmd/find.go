package cmd

import (
	"github.com/spf13/cobra"
	
	"github.com/ej-you/TemplateFilesManager/run_funcs"
)

// поиск файлов-шаблонов
var findCmd = &cobra.Command{
	Use:   "find",
	Short: "Find files-templates and its tags",
	Long: `Find files-templates and its tags

  Using «template find» will return all files-templates.
  Using «template find -m <match-string>» will return all files-templates which names containing substring <match-string>
  
  Using «template find -n <filename>» will return all tags of file-template with name <filename>
  Using «template find -n <filename> -m <match-string>» will return tags of file-template with name <filename> which names containing substring <match-string>`,
	RunE: run_funcs.FindRunE,
}


func init() {
	rootCmd.AddCommand(findCmd)

	// флаг для определения строки фильтра
	findCmd.Flags().StringP("match", "m", "", "Filter string (optional)")
	// флаг для определения имени файла
	findCmd.Flags().StringP("name", "n", "", "File-template name (optional)")
}

