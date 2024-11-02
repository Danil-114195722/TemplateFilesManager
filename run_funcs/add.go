package run_funcs

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	
	"github.com/spf13/cobra"

	"github.com/Danil-114195722/TemplateFilesManager/services"
	"github.com/Danil-114195722/TemplateFilesManager/settings"
)

func AddRunE(cmd *cobra.Command, args []string) error {
	// парсинг флагов имени и тега нового файла
	nameFlagValue, err := cmd.Flags().GetString("name")
	if err != nil {
		return err
	}
	tagFlagValue, err := cmd.Flags().GetString("tag")
	if err != nil {
		return err
	}

	// валидация имени файла
	match, _ := regexp.MatchString(`^[A-Za-zА-Яа-яЁё_\-\.\d]+$`, nameFlagValue)
	if !match {
		return errors.New("Invalid file name!")
	}

	isExists, err := services.FileTemplateIsExists(nameFlagValue, tagFlagValue)
	if err != nil {
		return err
	}
	// если такой файл с таким тегом уже есть в БД
	if isExists {
		settings.WarningPrintf("WARNING: file-template with such name and tag already exists!\n")
		settings.HintPrintf("\nHINT: use «template edit -n %s -t %s» to edit created file-template\n", nameFlagValue, tagFlagValue)
		return nil
	}

	// создание папки с названием файла в файловой системе
	err = os.Mkdir(fmt.Sprintf("%s/%s", settings.FilesPath, nameFlagValue), 0775)
	if err != nil && !os.IsExist(err) {
		return err
	}
	// создание файла с названием тега файла в файловой системе
	_, err = os.Create(fmt.Sprintf("%s/%s/%s", settings.FilesPath, nameFlagValue, tagFlagValue))
	if err != nil {
		return err
	}

	settings.SuccessPrintf("New file-template %q with tag %q created successfully!\n", nameFlagValue, tagFlagValue)
	settings.HintPrintf("\nHINT: use «template edit -n %s -t %s» to edit created file-template\n", nameFlagValue, tagFlagValue)
	return nil
}
