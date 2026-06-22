package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/rails_completer/cmd/common"
	"github.com/spf13/cobra"
)

var queryCmd = &cobra.Command{
	Use:   "query",
	Short: "Active Record read-only REPL",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(queryCmd).Standalone()

	queryCmd.Flags().String("database", "", "Database config name")
	queryCmd.Flags().String("db", "", "Alias for --database")
	queryCmd.Flags().Int("page", 1, "Page number")
	queryCmd.Flags().Int("per", 100, "Results per page (max 10000)")
	queryCmd.Flags().Bool("sql", false, "Treat input as raw SQL")

	common.AddEnvironmentFlag(queryCmd, "development")
}
