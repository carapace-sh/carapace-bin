package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Pull remote metadata and merge into local database",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pullCmd).Standalone()

	pullCmd.Flags().BoolP("help", "h", false, "Print help")
	pullCmd.Flags().BoolP("verbose", "v", false, "Show detailed information about pull decisions")
	rootCmd.AddCommand(pullCmd)

	carapace.Gen(pullCmd).PositionalCompletion(
		git.ActionRemotes(),
	)
}
