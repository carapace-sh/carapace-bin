package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/pub_completer/cmd/action"
	"github.com/spf13/cobra"
)

var global_deactivateCmd = &cobra.Command{
	Use:   "deactivate",
	Short: "Remove a previously activated package",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(global_deactivateCmd).Standalone()

	global_deactivateCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	globalCmd.AddCommand(global_deactivateCmd)

	carapace.Gen(global_deactivateCmd).PositionalCompletion(
		action.ActionActivePackages(),
	)
}
