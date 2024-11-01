package run_funcs

import (
	"errors"
	"fmt"
	"regexp"
	
	"github.com/spf13/cobra"

	"github.com/Danil-114195722/TemplateFilesManager/db/models"
	"github.com/Danil-114195722/TemplateFilesManager/db"
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

	// если такой файл с таким тегом уже есть в БД
	var file models.File
	dbConnect := db.GetConnection()
	selectResult := dbConnect.Where("name = ? AND tag = ?", nameFlagValue, tagFlagValue).First(&file)
	if selectResult.Error == nil { // NOT err
		return errors.New("File with such name and tag already exists!")
	}
	
	// создание новой записи файла в БД
	file.Name = nameFlagValue
	file.Tag = tagFlagValue
	createResult := dbConnect.Create(&file)
	if err = createResult.Error; err != nil {
		return err
	}

	if tagFlagValue == "default" {
		fmt.Printf("New file %q with default tag created successfully!\n", nameFlagValue)
		fmt.Printf("\nUse «template edit -n %s» to edit created file\n", nameFlagValue)
	} else {
		fmt.Printf("New file %q with tag %q created successfully!\n", nameFlagValue, tagFlagValue)
		fmt.Printf("\nUse «template edit -n %s -t %s» to edit created file\n", nameFlagValue, tagFlagValue)
	}
	return nil
}
