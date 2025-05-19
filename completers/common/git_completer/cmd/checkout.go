package cmd

import (
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var checkoutCmd = &cobra.Command{
	Use:     "checkout",
	Short:   "Switch branches or restore working tree files",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_main].ID,
}

func init() {
	carapace.Gen(checkoutCmd).Standalone()

	checkoutCmd.Flags().StringS("B", "B", "", "create/reset and checkout a branch")
	checkoutCmd.Flags().StringS("b", "b", "", "create and checkout a new branch")
	checkoutCmd.Flags().String("conflict", "", "conflict style (merge or diff3)")
	checkoutCmd.Flags().BoolP("detach", "d", false, "detach HEAD at named commit")
	checkoutCmd.Flags().BoolP("force", "f", false, "force checkout (throw away local modifications)")
	checkoutCmd.Flags().Bool("guess", false, "second guess 'git checkout <no-such-branch>' (default)")
	checkoutCmd.Flags().Bool("ignore-other-worktrees", false, "do not check if another worktree is holding the given ref")
	checkoutCmd.Flags().Bool("ignore-skip-worktree-bits", false, "do not limit pathspecs to sparse entries only")
	checkoutCmd.Flags().BoolS("l", "l", false, "create reflog for new branch")
	checkoutCmd.Flags().BoolP("merge", "m", false, "perform a 3-way merge with the new branch")
	checkoutCmd.Flags().Bool("no-guess", false, "do not second guess 'git checkout <no-such-branch>'")
	checkoutCmd.Flags().Bool("no-overlay", false, "remove files not in <tree-ish>")
	checkoutCmd.Flags().Bool("no-overwrite-ignore", false, "do not overwrite ignored files when switching branches")
	checkoutCmd.Flags().Bool("no-progress", false, "do not show progress")
	checkoutCmd.Flags().Bool("no-recurse-submodules", false, "do not update submodules")
	checkoutCmd.Flags().Bool("no-track", false, "do not set up \"upstream\" configuration")
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
		"B":                  git.ActionRefs(git.RefOption{LocalBranches: true}),
		"b":                  git.ActionRefs(git.RefOption{LocalBranches: true}),
		"conflict":           carapace.ActionValues("merge", "diff3"),
		"orphan":             git.ActionRefs(git.RefOption{LocalBranches: true, RemoteBranches: true}),
		"pathspec-from-file": carapace.ActionFiles(),
	})

	carapace.Gen(checkoutCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if strings.HasPrefix(c.Value, ".") {
				return git.ActionRefDiffs()
			}
			return git.ActionRefs(git.RefOption{}.Default())
		}),
	)

	carapace.Gen(checkoutCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			toFilter := make([]string, 0)
			for _, arg := range c.Args {
				// TODO directories? globs possible? should work for most cases though
				switch {
				case !strings.HasPrefix(c.Value, "."):
					toFilter = append(toFilter, strings.TrimPrefix(arg, "./"))
				case !strings.HasPrefix(arg, "."):
					toFilter = append(toFilter, "./"+arg)
				default:
					toFilter = append(toFilter, arg)
				}
			}

			// multipart completion here as there can be a lot of differences
			switch {
			case strings.HasPrefix(c.Args[0], "."):
				return git.ActionRefDiffs().Filter(toFilter...).MultiParts("/").StyleF(style.ForPathExt)
			default:
				return git.ActionRefDiffs(c.Args[0]).Filter(toFilter[1:]...).MultiParts("/").StyleF(style.ForPathExt)
			}
		}),
	)

	carapace.Gen(checkoutCmd).DashAnyCompletion(
		carapace.ActionPositional(checkoutCmd), // TODO should only be files, but will suffice for now,
	)
}
