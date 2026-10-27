package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var pullCmd = &cobra.Command{
	Use:     "pull",
	Short:   "Fetch from the remote and integrate it into the current branch",
	GroupID: "remote",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pullCmd).Standalone()

	pullCmd.Flags().Bool("all", false, "Fetch from every configured remote (respects `remote.<name>.skipFetchAll`)")
	pullCmd.Flags().Bool("autostash", false, "Stash dirty worktree/index changes around the pull (`merge.autoStash` / `rebase.autoStash`)")
	pullCmd.Flags().Bool("check", false, "Fetch and report whether a fast-forward is possible without changing the branch")
	pullCmd.Flags().String("deepen", "", "Deepen an existing shallow clone by this many commits (`git fetch --deepen`)")
	pullCmd.Flags().String("depth", "", "Fetch only the recent history of each branch (same as `git fetch --depth`)")
	pullCmd.Flags().Bool("dry-run", false, "Report ref updates without writing refs or objects")
	pullCmd.Flags().Bool("ff-only", false, "Refuse when a merge would be required (exit 1 when diverged)")
	pullCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	pullCmd.Flags().Bool("no-prune", false, "Do not prune stale remote-tracking refs, even when config enables it")
	pullCmd.Flags().Bool("no-rebase", false, "Merge instead of rebasing, overriding `pull.rebase`")
	pullCmd.Flags().Bool("no-tags", false, "Do not fetch tags automatically (`git fetch --no-tags`)")
	pullCmd.Flags().Bool("prune", false, "Remove remote-tracking refs that no longer exist on the remote")
	pullCmd.Flags().BoolP("quiet", "q", false, "Suppress progress and human output")
	pullCmd.Flags().Bool("rebase", false, "Rebase local commits onto the fetched upstream, overriding `pull.rebase=false`")
	pullCmd.Flags().Bool("tags", false, "Fetch all tags from the remote (`git fetch --tags`)")
	pullCmd.Flags().Bool("unshallow", false, "Convert a shallow repository to a complete one (`git fetch --unshallow`)")
	rootCmd.AddCommand(pullCmd)

	carapace.Gen(pullCmd).PositionalCompletion(
		git.ActionRemotes(),
	)
}
