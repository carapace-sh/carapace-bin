package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Pull, merge, rewrite if needed, and push metadata",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(syncCmd).Standalone()

	syncCmd.Flags().BoolP("help", "h", false, "Print help")
	syncCmd.Flags().BoolP("verbose", "v", false, "Show detailed information about sync decisions")
	rootCmd.AddCommand(syncCmd)

	carapace.Gen(syncCmd).PositionalCompletion(
		git.ActionRemotes(),
	)
}
