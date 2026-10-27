package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:     "show",
	Short:   "Show information about a commit, tag, or branch (defaults to HEAD)",
	GroupID: "history",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(showCmd).Standalone()

	showCmd.Flags().String("diff-merges", "combined", "Merge commit diffs: `combined` (default) or `first-parent`")
	showCmd.Flags().BoolP("function-context", "W", false, "Expand hunks to enclosing function boundaries (Git `-W` / `--function-context`)")
	showCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	showCmd.Flags().Bool("name-only", false, "Paths only, one per line")
	showCmd.Flags().Bool("name-status", false, "Status letter and path(s) per line")
	showCmd.Flags().Bool("numstat", false, "Machine-readable `insertions<TAB>deletions<TAB>path` lines")
	showCmd.Flags().BoolP("patch", "p", false, "Unified patch for each commit (`git log -p`)")
	showCmd.Flags().Bool("patch-text", false, "Git-apply-compatible unified patch (`git diff --binary` shape); not combinable with `--json`")
	showCmd.Flags().Bool("shortstat", false, "Only the summary line (`N files changed, …`)")
	showCmd.Flags().Bool("stat", false, "Summary with change bars per file (like `git diff --stat`)")
	showCmd.Flags().StringP("unified", "U", "", "Lines of context around each change (Git `-U`)")
	rootCmd.AddCommand(showCmd)

	carapace.Gen(showCmd).PositionalCompletion(
		git.ActionRefs(git.RefOption{}.Default()),
	)
}
