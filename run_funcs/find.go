package run_funcs

import (
	"fmt"
	
	"github.com/spf13/cobra"

	"github.com/ej-you/TemplateFilesManager/services"
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

	var foundResult []string
	var foundHeaderString string
	// если имя файла не указано, то получаем список всех найденнных файлов-шаблонов (с подстрокой match, если флаг был передан)
	if nameFlagValue == "" {
		foundResult, err = services.FindFilesTemplates(matchFlagValue)
		if err != nil {
			return err
		}
		foundHeaderString = "List of found files-templates:\n"
	// если имя файла было указано, то получаем список всех найденнных тегов (с подстрокой match, если флаг был передан) указанного файла-шаблона
	} else {
		foundResult, err = services.FindFileTemplateTags(nameFlagValue, matchFlagValue)
		if err != nil {
			return err
		}
		foundHeaderString = fmt.Sprintf("List of found tags of %q file-template:\n", nameFlagValue)
	}

	if len(foundResult) == 0 {
		fmt.Println("Nothing was found with given parameters")
		return nil
	}

	fmt.Println(foundHeaderString)
	// построчная печать найденных результатов
	for idx, foundEntity := range foundResult {
		fmt.Printf("  %d: %s\n", idx+1, foundEntity)
	}
	return nil
}
