package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/systemctl_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/ps"
	"github.com/spf13/cobra"
)

var helpCmd = &cobra.Command{
	Use:     "help",
	Short:   "Show manual for one or more units",
	GroupID: "unit",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(helpCmd).Standalone()

	rootCmd.AddCommand(helpCmd)

	carapace.Gen(helpCmd).PositionalAnyCompletion(
		carapace.Batch(
			action.ActionUnits(helpCmd),
			ps.ActionProcessIds(),
		).ToA().FilterArgs(),
	)
}
