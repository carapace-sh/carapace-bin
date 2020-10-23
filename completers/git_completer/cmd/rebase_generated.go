package cmd

import (
	"github.com/spf13/cobra"
)

var rebaseCmd = &cobra.Command{
	Use:   "rebase",
	Short: "Reapply commits on top of another base tip",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	rebaseCmd.Flags().Bool("abort", false, "abort and check out the original branch")
	rebaseCmd.Flags().Bool("apply", false, "use apply strategies to rebase")
	rebaseCmd.Flags().Bool("autosquash", false, "move commits that begin with squash!/fixup! under -i")
	rebaseCmd.Flags().Bool("autostash", false, "automatically stash/stash pop before and after")
	rebaseCmd.Flags().StringS("C", "C", "", "passed to 'git apply'")
	rebaseCmd.Flags().Bool("committer-date-is-author-date", false, "passed to 'git am'")
	rebaseCmd.Flags().Bool("continue", false, "continue")
	rebaseCmd.Flags().Bool("edit-todo", false, "edit the todo list during an interactive rebase")
	rebaseCmd.Flags().String("empty", "", "how to handle commits that become empty")
	rebaseCmd.Flags().BoolP("force-rebase", "f", false, "cherry-pick all commits, even if unchanged")
	rebaseCmd.Flags().Bool("fork-point", false, "use 'merge-base --fork-point' to refine upstream")
	rebaseCmd.Flags().Bool("ignore-date", false, "passed to 'git am'")
	rebaseCmd.Flags().Bool("ignore-whitespace", false, "passed to 'git am'")
	rebaseCmd.Flags().BoolP("interactive", "i", false, "let the user edit the list of commits to rebase")
	rebaseCmd.Flags().Bool("keep-base", false, "use the merge-base of upstream and branch as the current base")
	rebaseCmd.Flags().BoolP("merge", "m", false, "use merging strategies to rebase")
	rebaseCmd.Flags().BoolP("no-stat", "n", false, "do not show diffstat of what changed upstream")
	rebaseCmd.Flags().Bool("no-ff", false, "cherry-pick all commits, even if unchanged")
	rebaseCmd.Flags().Bool("no-verify", false, "allow pre-rebase hook to run")
	rebaseCmd.Flags().String("onto", "", "rebase onto given branch instead of upstream")
	rebaseCmd.Flags().BoolP("quiet", "q", false, "be quiet. implies --no-stat")
	rebaseCmd.Flags().Bool("quit", false, "abort but keep HEAD where it is")
	rebaseCmd.Flags().Bool("rerere-autoupdate", false, "update the index with reused conflict resolution if possible")
	rebaseCmd.Flags().Bool("reschedule-failed-exec", false, "automatically re-schedule any `exec` that fails")
	rebaseCmd.Flags().Bool("root", false, "rebase all reachable commits up to the root(s)")
	rebaseCmd.Flags().StringP("rebase-merges", "r", "", "try to rebase merges instead of skipping them")
	rebaseCmd.Flags().StringP("gpg-sign", "S", "", "GPG-sign commits")
	rebaseCmd.Flags().Bool("show-current-patch", false, "show the patch file being applied or merged")
	rebaseCmd.Flags().Bool("signoff", false, "add a Signed-off-by: line to each commit")
	rebaseCmd.Flags().Bool("skip", false, "skip current patch and continue")
	rebaseCmd.Flags().BoolP("strategy", "s", false, "<strategy>    use the given merge strategy")
	rebaseCmd.Flags().BoolP("verbose", "v", false, "display a diffstat of what changed upstream")
	rebaseCmd.Flags().String("whitespace", "", "passed to 'git apply'")
	rebaseCmd.Flags().BoolP("exec", "x", false, "<exec>     add exec lines after each commit of the editable list")
	rebaseCmd.Flags().BoolP("strategy-option", "X", false, "<option>    pass the argument through to the merge strategy")
	rootCmd.AddCommand(rebaseCmd)
}
