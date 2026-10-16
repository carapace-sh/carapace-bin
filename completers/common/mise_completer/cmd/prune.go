package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/mise_completer/cmd/action"
	"github.com/spf13/cobra"
)

var pruneCmd = &cobra.Command{
	Use:   "prune",
	Short: "Delete unused versions of tools",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pruneCmd).Standalone()

	pruneCmd.Flags().BoolP("dry-run", "n", false, "Show what would be removed without actually removing")
	pruneCmd.Flags().Bool("tools", false, "Only prune tool versions")
	rootCmd.AddCommand(pruneCmd)

	carapace.Gen(pruneCmd).PositionalAnyCompletion(
		action.ActionTools(),
	)
}
