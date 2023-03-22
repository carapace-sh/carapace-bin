package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/systemctl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var presetCmd = &cobra.Command{
	Use:     "preset",
	Short:   "Enable/disable one or more unit files based on preset configuration",
	GroupID: "unit file",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(presetCmd).Standalone()

	rootCmd.AddCommand(presetCmd)

	carapace.Gen(presetCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionUnits(presetCmd).Invoke(c).Filter(c.Args).ToA()
		}),
	)
}
