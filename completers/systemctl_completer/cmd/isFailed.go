package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/systemctl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var isFailedCmd = &cobra.Command{
	Use:     "is-failed",
	Short:   "Check whether units are failed",
	GroupID: "unit",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(isFailedCmd).Standalone()

	rootCmd.AddCommand(isFailedCmd)

	carapace.Gen(isFailedCmd).PositionalAnyCompletion(
		action.ActionUnits(isFailedCmd).FilterArgs(),
	)
}
