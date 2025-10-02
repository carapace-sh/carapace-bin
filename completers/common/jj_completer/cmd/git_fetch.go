package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/jj"
	"github.com/spf13/cobra"
)

var git_fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Fetch from a Git remote",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(git_fetchCmd).Standalone()

	git_fetchCmd.Flags().Bool("all-remotes", false, "Fetch from all remotes")
	git_fetchCmd.Flags().StringSlice("bookmark", nil, "Fetch only some of the branches")
	git_fetchCmd.Flags().StringSliceP("branch", "b", nil, "Fetch only some of the branches")
	git_fetchCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	git_fetchCmd.Flags().StringSlice("remote", []string{"origin"}, "The remote to fetch from (only named remotes are supported, can be repeated)")
	git_fetchCmd.Flags().Bool("tracked", false, "Fetch only tracked bookmarks")
	git_fetchCmd.Flag("bookmark").Hidden = true
	gitCmd.AddCommand(git_fetchCmd)

	actionBookmarks := carapace.ActionCallback(func(c carapace.Context) carapace.Action { // TODO verify
		remotes, err := git_fetchCmd.Flags().GetStringSlice("remote")
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		batch := carapace.Batch()
		for _, remote := range remotes {
			batch = append(batch, jj.ActionRemoteBookmarks(remote))
		}
		return batch.ToA()
	})

	carapace.Gen(git_fetchCmd).FlagCompletion(carapace.ActionMap{
		"bookmark": actionBookmarks,
		"branch":   actionBookmarks,
		"remote":   jj.ActionRemotes(),
	})
}
