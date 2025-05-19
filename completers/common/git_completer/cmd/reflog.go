package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var reflogCmd = &cobra.Command{
	Use:                "reflog",
	Short:              "Manage reflog information",
	Run:                func(cmd *cobra.Command, args []string) {},
	GroupID:            groups[group_manipulator].ID,
	DisableFlagParsing: true,
}

func init() {
	carapace.Gen(reflogCmd).Standalone()

	rootCmd.AddCommand(reflogCmd)

	carapace.Gen(reflogCmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin("git", "reflog", "show"),
	)
}
