package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/systemctl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var addRequiresCmd = &cobra.Command{
	Use:     "add-requires",
	Short:   "Add 'Requires' dependency for the target on specified one or more units",
	GroupID: "unit file",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(addRequiresCmd).Standalone()

	rootCmd.AddCommand(addRequiresCmd)

	carapace.Gen(addRequiresCmd).PositionalCompletion(
		action.ActionTargets(addRequiresCmd),
	)

	carapace.Gen(addRequiresCmd).PositionalAnyCompletion(
		action.ActionUnits(addRequiresCmd).FilterArgs(),
	)
}
