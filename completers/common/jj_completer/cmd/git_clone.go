package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var git_cloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "Create a new repo backed by a clone of a Git repo",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(git_cloneCmd).Standalone()

	git_cloneCmd.Flags().StringSlice("bookmark", nil, "Name of the branch to fetch and use as the parent of the working-copy change (can be repeated)")
	git_cloneCmd.Flags().StringSliceP("branch", "b", nil, "Name of the branch to fetch and use as the parent of the working-copy change (can be repeated)")
	git_cloneCmd.Flags().Bool("colocate", false, "Colocate the Jujutsu repo with the git repo")
	git_cloneCmd.Flags().String("depth", "", "Create a shallow clone of the given depth")
	git_cloneCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	git_cloneCmd.Flags().Bool("no-colocate", false, "Disable colocation of the Jujutsu repo with the git repo")
	git_cloneCmd.Flags().String("object-hash", "", "Object hash algorithm for the local Git repository")
	git_cloneCmd.Flags().String("remote", "origin", "Name of the newly created remote")
	git_cloneCmd.Flags().StringSliceP("tag", "t", nil, "Fetch only some of the tags (can be repeated)")
	git_cloneCmd.Flag("bookmark").Hidden = true
	gitCmd.AddCommand(git_cloneCmd)

	carapace.Gen(git_cloneCmd).FlagCompletion(carapace.ActionMap{
		"bookmark": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) == 0 {
				return carapace.ActionValues()
			}
			return git.ActionLsRemoteRefs(git.LsRemoteRefOption{Url: c.Args[0], Branches: true})
		}),
		"branch": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) == 0 {
				return carapace.ActionValues()
			}
			return git.ActionLsRemoteRefs(git.LsRemoteRefOption{Url: c.Args[0], Branches: true})
		}),
		"object-hash": carapace.ActionValues("sha1", "sha256"),
	})

	carapace.Gen(git_cloneCmd).PositionalCompletion(
		git.ActionRepositorySearch(git.SearchOpts{}.Default()),
		carapace.ActionDirectories(),
	)
}
