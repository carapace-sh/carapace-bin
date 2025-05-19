package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/waypoint_completer/cmd/action"
	"github.com/spf13/cobra"
)

var authMethod_inspectCmd = &cobra.Command{
	Use:   "inspect",
	Short: "Show detailed information about a configured auth method",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(authMethod_inspectCmd).Standalone()

	authMethodCmd.AddCommand(authMethod_inspectCmd)

	carapace.Gen(authMethod_inspectCmd).PositionalCompletion(
		action.ActionAuthMethods(),
	)
}
