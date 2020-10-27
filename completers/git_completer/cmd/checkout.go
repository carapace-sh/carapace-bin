package cmd

import (
	"os"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/git"
	"github.com/spf13/cobra"
)

var checkoutCmd = &cobra.Command{
	Use:   "checkout",
	Short: "Switch branches or restore working tree files",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(checkoutCmd).Standalone()

	checkoutCmd.Flags().StringS("b", "b", "", "create and checkout a new branch")
	checkoutCmd.Flags().StringS("B", "B", "", "create/reset and checkout a branch")
	checkoutCmd.Flags().BoolS("l", "l", false, "create reflog for new branch")
	checkoutCmd.Flags().String("conflict", "", "conflict style (merge or diff3)")
	checkoutCmd.Flags().BoolP("detach", "d", false, "detach HEAD at named commit")
	checkoutCmd.Flags().BoolP("force", "f", false, "force checkout (throw away local modifications)")
	checkoutCmd.Flags().Bool("guess", false, "second guess 'git checkout <no-such-branch>' (default)")
	checkoutCmd.Flags().Bool("ignore-other-worktrees", false, "do not check if another worktree is holding the given ref")
	checkoutCmd.Flags().Bool("ignore-skip-worktree-bits", false, "do not limit pathspecs to sparse entries only")
	checkoutCmd.Flags().BoolP("merge", "m", false, "perform a 3-way merge with the new branch")
	checkoutCmd.Flags().String("orphan", "", "new unparented branch")
	checkoutCmd.Flags().BoolP("ours", "2", false, "checkout our version for unmerged files")
	checkoutCmd.Flags().Bool("overlay", false, "use overlay mode (default)")
	checkoutCmd.Flags().Bool("overwrite-ignore", false, "update ignored files (default)")
	checkoutCmd.Flags().BoolP("patch", "p", false, "select hunks interactively")
	checkoutCmd.Flags().Bool("pathspec-file-nul", false, "pathspec elements are separated with NUL character")
	checkoutCmd.Flags().String("pathspec-from-file", "", "read pathspec from file")
	checkoutCmd.Flags().Bool("progress", false, "force progress reporting")
	checkoutCmd.Flags().BoolP("quiet", "q", false, "suppress progress reporting")
	checkoutCmd.Flags().String("recurse-submodules", "", "control recursive updating of submodules")
	checkoutCmd.Flags().BoolP("theirs", "3", false, "checkout their version for unmerged files")
	checkoutCmd.Flags().BoolP("track", "t", false, "set upstream info for new branch")
	rootCmd.AddCommand(checkoutCmd)

	carapace.Gen(checkoutCmd).FlagCompletion(carapace.ActionMap{
		"b":                  git.ActionRefs(git.RefOption{LocalBranches: true}),
		"B":                  git.ActionRefs(git.RefOption{LocalBranches: true}),
		"conflict":           carapace.ActionValues("merge", "diff3"),
		"orphan":             git.ActionRefs(git.RefOption{LocalBranches: true, RemoteBranches: true}),
		"pathspec-from-file": carapace.ActionFiles(""),
	})

	carapace.Gen(checkoutCmd).PositionalCompletion(
		git.ActionRefs(git.RefOptionDefault),
	)

	carapace.Gen(checkoutCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(args []string) carapace.Action {
			// the first `--` is currently not detected as positional argument,
			// so just search the full command line for it and assume it's the divider
			for _, arg := range os.Args {
				if arg == "--" {
					return carapace.ActionFiles("") // TODO files from branch?
				}
			}
			return carapace.ActionValues()
		}),
	)
}
