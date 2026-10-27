package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var pushCmd = &cobra.Command{
	Use:     "push",
	Short:   "Publish refs to a remote",
	GroupID: "remotes",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pushCmd).Standalone()

	pushCmd.Flags().Bool("atomic", false, "All ref updates succeed or none are applied")
	pushCmd.Flags().StringSlice("delete", nil, "Delete a remote branch or tag (may be repeated)")
	pushCmd.Flags().Bool("dry-run", false, "Contact the remote but do not update refs or send objects")
	pushCmd.Flags().Bool("follow-tags", false, "Push annotated tags reachable from the commits being pushed")
	pushCmd.Flags().BoolP("force", "f", false, "Overwrite remote refs even when the update is not a fast-forward")
	pushCmd.Flags().String("force-with-lease", "", "Push only when the remote ref matches the local remote-tracking ref, or `=<rev>` when pushing exactly one ref")
	pushCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	pushCmd.Flags().StringSliceP("push-option", "o", nil, "Server-side push option (repeatable; `GIT_PUSH_OPTION_*` in hooks)")
	pushCmd.Flags().String("remote", "", "Remote to push to (defaults to the current branch's remote or `origin`)")
	pushCmd.Flags().BoolP("set-upstream", "u", false, "Set `branch.<name>.remote` and `branch.<name>.merge` for each pushed branch")
	pushCmd.Flags().BoolP("tags", "t", false, "Push tags instead of the current branch. Pushes every local tag under `refs/tags/` to the remote")
	pushCmd.Flag("force-with-lease").NoOptDefVal = " "
	pushCmd.Flag("remote").Hidden = true
	rootCmd.AddCommand(pushCmd)

	carapace.Gen(pushCmd).FlagCompletion(carapace.ActionMap{
		"delete":           git.ActionRefs(git.RefOption{LocalBranches: true, RemoteBranches: true, Tags: true}),
		"force-with-lease": git.ActionRefs(git.RefOption{}.Default()),
		"remote":           git.ActionRemotes(),
	})

	carapace.Gen(pushCmd).PositionalCompletion(
		git.ActionRemotes(),
	)
}
