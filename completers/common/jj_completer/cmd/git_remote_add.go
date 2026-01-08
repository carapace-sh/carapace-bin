package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/jj"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var git_remote_addCmd = &cobra.Command{
	Use:   "add [OPTIONS] <REMOTE> <URL>",
	Short: "Add a Git remote",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(git_remote_addCmd).Standalone()

	git_remote_addCmd.Flags().String("fetch-tags", "", "Configure when to fetch tags")
	git_remote_addCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	git_remote_addCmd.Flags().String("push-url", "", "The URL used for push")
	git_remoteCmd.AddCommand(git_remote_addCmd)

	carapace.Gen(git_remote_addCmd).FlagCompletion(carapace.ActionMap{
		"fetch-tags": carapace.ActionValues("all", "included", "none").StyleF(style.ForKeyword),
	})

	carapace.Gen(git_remote_addCmd).PositionalCompletion(
		jj.ActionRemotes(),
		git.ActionRepositorySearch(git.SearchOpts{}.Default()),
	)
}
