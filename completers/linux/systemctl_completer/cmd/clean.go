package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/systemctl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var cleanCmd = &cobra.Command{
	Use:     "clean",
	Short:   "Clean runtime, cache, state, logs or configuration of unit",
	GroupID: "unit",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanCmd).Standalone()

	rootCmd.AddCommand(cleanCmd)

	carapace.Gen(cleanCmd).PositionalAnyCompletion(
		action.ActionUnits(cleanCmd).FilterArgs(),
	)
}
