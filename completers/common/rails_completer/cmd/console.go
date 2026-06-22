package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/rails_completer/cmd/common"
	"github.com/spf13/cobra"
)

var consoleCmd = &cobra.Command{
	Use:   "console",
	Short: "Start the Rails console",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(consoleCmd).Standalone()

	consoleCmd.Flags().BoolP("query-cache", "q", false, "Enable query cache")
	consoleCmd.Flags().BoolP("sandbox", "s", false, "Rollback database changes on exit")
	consoleCmd.Flags().BoolP("skip-executor", "w", false, "Don't wrap with Rails Executor")

	common.AddEnvironmentFlag(consoleCmd, "development")
}
