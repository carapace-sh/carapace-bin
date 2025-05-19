package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/systemctl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:     "status",
	Short:   "Show runtime status of one or more units",
	GroupID: "unit",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(statusCmd).Standalone()

	rootCmd.AddCommand(statusCmd)

	carapace.Gen(statusCmd).PositionalAnyCompletion(
		action.ActionUnits(statusCmd).FilterArgs(),
	)
}
