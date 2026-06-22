package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/rails_completer/cmd/common"
	"github.com/spf13/cobra"
)

var runnerCmd = &cobra.Command{
	Use:   "runner",
	Short: "Execute Ruby code in the Rails application context",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runnerCmd).Standalone()

	runnerCmd.Flags().BoolP("skip-executor", "w", false, "Don't wrap with Rails Executor")

	common.AddEnvironmentFlag(runnerCmd, "development")

	carapace.Gen(runnerCmd).PositionalCompletion(
		carapace.ActionFiles("*.rb"),
	)
}
