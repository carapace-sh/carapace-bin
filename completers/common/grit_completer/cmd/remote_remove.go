package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var remote_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a remote and its remote-tracking refs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(remote_removeCmd).Standalone()

	remote_removeCmd.Flags().Bool("dry-run", false, "Report changes without writing config or refs")
	remote_removeCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	remoteCmd.AddCommand(remote_removeCmd)

	carapace.Gen(remote_removeCmd).PositionalCompletion(
		git.ActionRemotes(),
	)
}
