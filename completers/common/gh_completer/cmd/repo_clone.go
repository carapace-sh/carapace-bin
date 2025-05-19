package cmd

import (
	"fmt"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/gh_completer/cmd/action"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var repo_cloneCmd = &cobra.Command{
	Use:     "clone <repository> [<directory>] [-- <gitflags>...]",
	Short:   "Clone a repository locally",
	GroupID: "Targeted commands",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repo_cloneCmd).Standalone()

	repo_cloneCmd.Flags().StringP("upstream-remote-name", "u", "", "Upstream remote name when cloning a fork")
	repoCmd.AddCommand(repo_cloneCmd)

	carapace.Gen(repo_cloneCmd).PositionalCompletion(
		action.ActionOwnerRepositories(repo_cloneCmd),
		carapace.ActionDirectories(),
	)

	carapace.Gen(repo_cloneCmd).DashAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			repo := ""
			if args := repo_cloneCmd.Flags().Args(); len(args) > 0 {
				repo = fmt.Sprintf("https://github.com/%v.git", args[0])
			}
			return bridge.ActionCarapaceBin("git", "clone", repo, "")
		}),
	)
}
