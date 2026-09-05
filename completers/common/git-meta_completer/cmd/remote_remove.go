package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var remote_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a metadata remote source",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(remote_removeCmd).Standalone()

	remote_removeCmd.Flags().BoolP("help", "h", false, "Print help")
	remoteCmd.AddCommand(remote_removeCmd)

	carapace.Gen(remote_removeCmd).PositionalCompletion(
		git.ActionRemotes(),
	)
}
