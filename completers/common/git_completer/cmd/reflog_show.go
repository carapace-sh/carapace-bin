package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var reflog_showCmd = &cobra.Command{
	Use:                "show",
	Short:              "Show the log",
	Run:                func(cmd *cobra.Command, args []string) {},
	DisableFlagParsing: true,
}

func init() {
	carapace.Gen(reflog_showCmd).Standalone()

	reflogCmd.AddCommand(reflog_showCmd)

	carapace.Gen(reflog_showCmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin("git", "log", "-g", "--abbrev-commit", "--pretty=oneline"),
	)
}
