package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/systemctl_completer/cmd/action"
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
		action.ActionUnits(linkCmd).FilterArgs(),
	)
}
