package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/systemctl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var catCmd = &cobra.Command{
	Use:     "cat",
	Short:   "Show files and drop-ins of specified units",
	GroupID: "unit",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(catCmd).Standalone()

	rootCmd.AddCommand(catCmd)

	carapace.Gen(catCmd).PositionalAnyCompletion(
		action.ActionUnits(catCmd).FilterArgs(),
	)
}
