package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/systemctl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var linkCmd = &cobra.Command{
	Use:     "link",
	Short:   "Link one or more units files into the search path",
	GroupID: "unit file",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(linkCmd).Standalone()

	rootCmd.AddCommand(linkCmd)

	carapace.Gen(linkCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionUnits(linkCmd).Invoke(c).Filter(c.Args).ToA()
		}),
	)
}
