package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/systemctl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var unmaskCmd = &cobra.Command{
	Use:     "unmask",
	Short:   "Unmask one or more units",
	GroupID: "unit file",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(unmaskCmd).Standalone()

	rootCmd.AddCommand(unmaskCmd)

	carapace.Gen(unmaskCmd).PositionalAnyCompletion(
		action.ActionUnits(unmaskCmd).FilterArgs(),
	)
}
