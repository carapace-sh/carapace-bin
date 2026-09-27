package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/asdf_completer/cmd/action"
	"github.com/spf13/cobra"
)

var latestCmd = &cobra.Command{
	Use:   "latest",
	Short: "Show latest stable version of a package",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(latestCmd).Standalone()

	latestCmd.Flags().Bool("all", false, "Show latest version of all tools")
	rootCmd.AddCommand(latestCmd)

	carapace.Gen(latestCmd).PositionalCompletion(
		action.ActionPlugins(),
	)
}
