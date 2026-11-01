package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Push local metadata to a remote",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pushCmd).Standalone()

	pushCmd.Flags().BoolP("help", "h", false, "Print help")
	pushCmd.Flags().Bool("readme", false, "Push a README to refs/heads/main on the meta remote (only if it doesn't already exist)")
	pushCmd.Flags().BoolP("verbose", "v", false, "Show detailed information about push decisions")
	rootCmd.AddCommand(pushCmd)

	carapace.Gen(pushCmd).PositionalCompletion(
		git.ActionRemotes(),
	)
}
