package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/git"
	"github.com/spf13/cobra"
)

var remote_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a remote",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(remote_addCmd).Standalone()

	remote_addCmd.Flags().BoolP("fetch", "f", false, "fetch the remote branches")
	remote_addCmd.Flags().StringP("master", "m", "", "master branch")
	remote_addCmd.Flags().String("mirror", "", "set up remote as a mirror to push to or fetch from")
	remote_addCmd.Flags().Bool("tags", false, "import all tags and associated objects when fetching")
	remote_addCmd.Flags().StringArrayP("track", "t", []string{}, "branch(es) to track")

	remote_addCmd.Flag("mirror").NoOptDefVal = " "

	remoteCmd.AddCommand(remote_addCmd)

	carapace.Gen(remote_addCmd).FlagCompletion(carapace.ActionMap{
		"mirror": carapace.ActionValues("push", "fetch"),
		"master": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return git.ActionLsRemoteRefs(c.Args[1], git.LsRemoteRefOption{Branches: true})
		}),
		"track": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return git.ActionLsRemoteRefs(c.Args[1], git.LsRemoteRefOption{Branches: true})
		}),
	})

	carapace.Gen(remote_addCmd).PositionalCompletion(
		git.ActionRemotes(),
	)
}
