package cmd

import (
	"github.com/spf13/cobra"
	
	"github.com/ej-you/TemplateFilesManager/run_funcs"
)

// manageCmd represents the manage command
var manageCmd = &cobra.Command{
	Use:   "manage",
	Short: "Provides manage.sh funcs like \"status\" and \"uninstall\"",
	Long: `Provides manage.sh funcs like status and uninstall
  Use «template manage» to get help from manage.sh`,
	RunE: run_funcs.ManageRunE,
}

func init() {
	rootCmd.AddCommand(manageCmd)
}
