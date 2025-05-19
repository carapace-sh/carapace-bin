package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/systemctl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var listSocketsCmd = &cobra.Command{
	Use:     "list-sockets",
	Short:   "List socket units currently in memory, ordered by address",
	GroupID: "unit",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listSocketsCmd).Standalone()

	rootCmd.AddCommand(listSocketsCmd)

	carapace.Gen(listSocketsCmd).PositionalAnyCompletion(
		action.ActionUnits(listSocketsCmd).FilterArgs(),
	)
}
