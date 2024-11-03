package run_funcs

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"

	"github.com/Danil-114195722/TemplateFilesManager/settings"
)


func ManageRunE(cmd *cobra.Command, args []string) error {
	// если команда на удаление утилиты, то запрашиваем подтверждение
	if len(args) > 0 && args[0] == "uninstall" {
		var confirmation string
		fmt.Print("You really want to uninstall template utility? [y, n] ")
		fmt.Scan(&confirmation)

		if confirmation != "y" {
			fmt.Println("Canceled")
			return nil
		}
	}

	// создание команды для запуска скрипта manage.sh
	command := exec.Command(fmt.Sprintf("%s/manage.sh", settings.BaseDir), args...)
	
	// для корректной передачи вывода из запускаемого скрипта в утилиту
	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	// выполнение команды
	err := command.Run()
	if err != nil {
		return err
	}
	return nil
}
