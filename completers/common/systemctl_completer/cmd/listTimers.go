package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/systemctl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var listTimersCmd = &cobra.Command{
	Use:     "list-timers",
	Short:   "List timer units currently in memory, ordered by next elapse",
	GroupID: "unit",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listTimersCmd).Standalone()

	rootCmd.AddCommand(listTimersCmd)

	carapace.Gen(listTimersCmd).PositionalAnyCompletion(
		action.ActionUnits(listTimersCmd).FilterArgs(),
	)
}
