package run_funcs

import (
	"fmt"
	
	"github.com/spf13/cobra"

	"github.com/ej-you/TemplateFilesManager/services"
	"github.com/ej-you/TemplateFilesManager/settings"
)

func DeleteRunE(cmd *cobra.Command, args []string) error {
	// парсинг флагов имени и тега файла и флага на удаление всех файлов-шаблонов (всех тегов одного файла шаблона)
	nameFlagValue, err := cmd.Flags().GetString("name")
	if err != nil {
		return err
	}
	tagFlagValue, err := cmd.Flags().GetString("tag")
	if err != nil {
		return err
	}
	allFlagValue, err := cmd.Flags().GetBool("all")
	if err != nil {
		return err
	}

	var confirmation string
	// удаление всех файлов-шаблонов
	if allFlagValue && nameFlagValue == "" {
		// запрашиваем подтверждение на удаление
		fmt.Print("Are you sure you want to delete ALL files-templates? [y, n] ")
		fmt.Scan(&confirmation)
		if confirmation != "y" {
			fmt.Println("Canceled")
			return nil
		}
		fmt.Println("Continue...")

		// удаление всех директорий файлов-шаблонов
		successDeletedCount, err := services.DeleteAllFilesTemplates()
		// если ни один файл-шаблон не найден
		if err != nil && err.Error() == "files-templates not found" {
			settings.WarningPrintf("WARNING: no files-templates found!\n")
			fmt.Println("Nothing was removed")
			return nil
		}
		if err != nil {
			return err
		}
		settings.SuccessPrintf("Successfully removed all (%d) files-templates\n", successDeletedCount)

	// удаление всех тегов определённого файла-шаблона
	} else if allFlagValue && nameFlagValue != "" {
		// запрашиваем подтверждение на удаление
		fmt.Printf("Are you sure you want to delete all tags of %q file-template? [y, n] ", nameFlagValue)
		fmt.Scan(&confirmation)
		if confirmation != "y" {
			fmt.Println("Canceled")
			return nil
		}
		fmt.Println("Continue...")

		// если такого файла-шаблона нет в файловой системе
		isExists, err := services.FileTemplateIsExists(nameFlagValue)
		if err != nil {
			return err
		}
		if !isExists {
			settings.WarningPrintf("WARNING: file-template with such name does not exist!\n")
			fmt.Println("Nothing was removed")
			return nil
		}

		// удаление директории файла-шаблона
		err = services.DeleteFileTemplate(nameFlagValue)
		if err != nil {
			return err
		}
		settings.SuccessPrintf("Successfully removed 1 file-template\n")

	// удаление определённого тега определённого файла-шаблона
	} else {
		// запрашиваем подтверждение на удаление
		fmt.Printf("Are you sure you want to delete tag %q of %q file-template? [y, n] ", tagFlagValue, nameFlagValue)
		fmt.Scan(&confirmation)
		if confirmation != "y" {
			fmt.Println("Canceled")
			return nil
		}
		fmt.Println("Continue...")

		// если такого файла с таким тегом нет в файловой системе
		isExists, err := services.FileTemplateTagIsExists(nameFlagValue, tagFlagValue)
		if err != nil {
			return err
		}
		if !isExists {
			settings.WarningPrintf("WARNING: file-template with such name and tag does not exist!\n")
			fmt.Println("Nothing was removed")
			return nil
		}

		// удаление файла тега
		err = services.DeleteFileTemplateTag(nameFlagValue, tagFlagValue)
		if err != nil {
			return err
		}
		settings.SuccessPrintf("Successfully removed 1 tag of file-template\n")
	}

	return nil
}
