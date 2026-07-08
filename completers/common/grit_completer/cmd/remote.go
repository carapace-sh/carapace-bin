package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var remoteCmd = &cobra.Command{
	Use:   "remote",
	Short: "List remotes, or add one",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(remoteCmd).Standalone()

	remoteCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.AddCommand(remoteCmd)

	carapace.Gen(remoteCmd).PositionalAnyCompletion(
		git.ActionRemotes(),
	)
}
