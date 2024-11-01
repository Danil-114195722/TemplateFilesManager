package run_funcs

import (
	"fmt"
	
	"github.com/spf13/cobra"

	"github.com/Danil-114195722/TemplateFilesManager/services"
	// "github.com/Danil-114195722/TemplateFilesManager/settings"
)


func FindRunE(cmd *cobra.Command, args []string) error {
	// парсинг флагов имени файла и подстроки для поиска
	nameFlagValue, err := cmd.Flags().GetString("name")
	if err != nil {
		return err
	}
	matchFlagValue, err := cmd.Flags().GetString("match")
	if err != nil {
		return err
	}

	// если имя файла не указано, то выдаём список всех найденнных файлов-шаблонов (с подстрокой match, если флаг был передан)
	if nameFlagValue == "" {
		foundFiles, err := services.FindFilesTemplates(matchFlagValue)
		if err != nil {
			return err
		}
		fmt.Println("foundFiles:", foundFiles)
		return nil
	}
	// если имя файла было указано, то выдаём список всех найденнных тегов (с подстрокой match, если флаг был передан) указанного файла-шаблона
	foundFileTags, err := services.FindFileTemplateTags(nameFlagValue, matchFlagValue)
	if err != nil {
		return err
	}
	fmt.Println("foundFileTags:", foundFileTags)
	return nil
}