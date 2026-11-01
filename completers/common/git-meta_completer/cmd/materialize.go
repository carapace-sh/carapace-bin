package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var materializeCmd = &cobra.Command{
	Use:   "materialize",
	Short: "Materialize remote metadata into local SQLite",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(materializeCmd).Standalone()

	materializeCmd.Flags().Bool("dry-run", false, "Show what would be changed without updating SQLite or refs")
	materializeCmd.Flags().Bool("force-full", false, "Reindex promisor keys from the full remote metadata history")
	materializeCmd.Flags().BoolP("help", "h", false, "Print help")
	materializeCmd.Flags().BoolP("verbose", "v", false, "Show detailed information about merge decisions and tree parsing")
	rootCmd.AddCommand(materializeCmd)

	carapace.Gen(materializeCmd).PositionalCompletion(
		git.ActionRemotes(),
	)
}
