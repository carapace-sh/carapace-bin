package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/rails_completer/cmd/common"
	"github.com/spf13/cobra"
)

var dbconsoleCmd = &cobra.Command{
	Use:   "dbconsole",
	Short: "Start a console for the database",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dbconsoleCmd).Standalone()

	dbconsoleCmd.Flags().String("database", "", "Database config name")
	dbconsoleCmd.Flags().String("db", "", "Alias for --database")
	dbconsoleCmd.Flags().Bool("header", false, "Show headers")
	dbconsoleCmd.Flags().BoolP("include-password", "p", false, "Auto-provide password")
	dbconsoleCmd.Flags().String("mode", "", "SQLite3 mode (html, list, line, column)")

	common.AddEnvironmentFlag(dbconsoleCmd, "development")

	carapace.Gen(dbconsoleCmd).FlagCompletion(carapace.ActionMap{
		"mode": carapace.ActionValues("html", "list", "line", "column"),
	})
}
