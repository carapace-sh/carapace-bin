package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/jj"
	"github.com/spf13/cobra"
)

var git_remote_removeCmd = &cobra.Command{
	Use:   "remove [OPTIONS] <REMOTE>",
	Short: "Remove a Git remote and forget its bookmarks",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(git_remote_removeCmd).Standalone()

	git_remote_removeCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	git_remoteCmd.AddCommand(git_remote_removeCmd)

	carapace.Gen(git_remote_removeCmd).PositionalCompletion(
		jj.ActionRemotes(),
	)
}
