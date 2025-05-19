package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/npm_completer/cmd/action"
	"github.com/spf13/cobra"
)

var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "View registry info",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(viewCmd).Standalone()
	viewCmd.Flags().Bool("json", false, "output as json")
	addWorkspaceFlags(viewCmd)

	rootCmd.AddCommand(viewCmd)

	carapace.Gen(viewCmd).PositionalCompletion(
		action.ActionPackages(viewCmd),
		// TODO field completion
	)
}
