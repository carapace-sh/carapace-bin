package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var remote_setUrlCmd = &cobra.Command{
	Use:   "set-url",
	Short: "Set the fetch or push URL for an existing remote",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(remote_setUrlCmd).Standalone()

	remote_setUrlCmd.Flags().Bool("dry-run", false, "Show what would change without writing config")
	remote_setUrlCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	remote_setUrlCmd.Flags().Bool("push", false, "Update `pushurl` instead of `url`")
	remoteCmd.AddCommand(remote_setUrlCmd)

	carapace.Gen(remote_setUrlCmd).PositionalCompletion(
		git.ActionRemotes(),
		git.ActionRepositorySearch(git.SearchOpts{}.Default()),
	)
}
