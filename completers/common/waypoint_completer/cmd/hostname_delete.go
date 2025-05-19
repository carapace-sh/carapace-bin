package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/waypoint_completer/cmd/action"
	"github.com/spf13/cobra"
)

var hostname_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a previously registered hostname",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(hostname_deleteCmd).Standalone()

	addGlobalOptions(hostname_deleteCmd)

	hostnameCmd.AddCommand(hostname_deleteCmd)

	carapace.Gen(hostname_deleteCmd).PositionalCompletion(
		action.ActionHostnames(),
	)
}
