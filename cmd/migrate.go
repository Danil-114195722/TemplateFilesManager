package cmd

import (
	"github.com/spf13/cobra"

	"github.com/Danil-114195722/TemplateFilesManager/db"
)


// миграция БД (скрытая команда)
var migrateCmd = &cobra.Command{
	Use:	"migrate",
	Long:	`Migrate SQLite DB.`,
	Hidden: true,
	Run: func(cmd *cobra.Command, args []string) {
		db.Migrate()
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}
