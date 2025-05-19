package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/jj"
	"github.com/spf13/cobra"
)

var git_remote_setUrlCmd = &cobra.Command{
	Use:   "set-url [OPTIONS] <REMOTE> <URL>",
	Short: "Set the URL of a Git remote",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(git_remote_setUrlCmd).Standalone()

	git_remote_setUrlCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	git_remoteCmd.AddCommand(git_remote_setUrlCmd)

	carapace.Gen(git_remote_setUrlCmd).PositionalCompletion(
		jj.ActionRemotes(),
		git.ActionRepositorySearch(git.SearchOpts{}.Default()),
	)
}
