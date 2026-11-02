package cmd

import (
	"fmt"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
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

	repo_cloneCmd.Flags().Bool("no-upstream", false, "Do not add an upstream remote when cloning a fork")
	repo_cloneCmd.Flags().StringP("upstream-remote-name", "u", "upstream", "Upstream remote name when cloning a fork")
	repoCmd.AddCommand(repo_cloneCmd)

	carapace.Gen(repo_cloneCmd).PositionalCompletion(
		gh.ActionOwnerRepositories(gh.HostOpts{}),
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
