package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Fetch from the remote and integrate it into the current branch",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pullCmd).Standalone()

	pullCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.AddCommand(pullCmd)

	carapace.Gen(pullCmd).PositionalCompletion(
		git.ActionRemotes(),
	)
}
