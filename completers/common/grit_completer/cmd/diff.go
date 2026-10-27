package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var diffCmd = &cobra.Command{
	Use:     "diff",
	Short:   "Show changes as a diff. No argument: uncommitted changes; with a commit: the change that commit introduced; with two commits (or `A..B`) diff their trees",
	GroupID: "changes",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(diffCmd).Standalone()

	diffCmd.Flags().String("against", "", "Compare the worktree to this ref's tree instead of HEAD (no commit arguments)")
	diffCmd.Flags().Bool("cached", false, "Staged changes only (index vs HEAD)")
	diffCmd.Flags().String("diff-algorithm", "", "Line diff algorithm (overrides `diff.algorithm`)")
	diffCmd.Flags().String("diff-filter", "", "Limit output to these change types (Git letter syntax: `ADMR…`)")
	diffCmd.Flags().Bool("exit-code", false, "Exit with 1 when there are changes (errors exit 2)")
	diffCmd.Flags().StringP("find-renames", "M", "", "Detect renames; optional similarity threshold (`50` default, or `50%`). Use `-M`, `-M30`, or `--find-renames` / `--find-renames=30` (long form requires `=` when a threshold is given)")
	diffCmd.Flags().BoolP("function-context", "W", false, "Expand hunks to enclosing function boundaries (Git `-W` / `--function-context`)")
	diffCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	diffCmd.Flags().BoolP("ignore-all-space", "w", false, "Ignore whitespace when comparing lines (`git diff -w`)")
	diffCmd.Flags().Bool("ignore-blank-lines", false, "Ignore changes whose lines are all blank (`git diff --ignore-blank-lines`)")
	diffCmd.Flags().BoolP("ignore-space-change", "b", false, "Ignore changes in the amount of whitespace (`git diff -b`)")
	diffCmd.Flags().Bool("name-only", false, "Paths only, one per line")
	diffCmd.Flags().Bool("name-status", false, "Status letter and path(s) per line")
	diffCmd.Flags().Bool("no-index", false, "Compare two filesystem paths outside any repository (like `git diff --no-index`)")
	diffCmd.Flags().Bool("no-renames", false, "Do not detect renames")
	diffCmd.Flags().Bool("numstat", false, "Machine-readable `insertions<TAB>deletions<TAB>path` lines")
	diffCmd.Flags().BoolP("patch", "p", false, "Unified patch for each commit (`git log -p`)")
	diffCmd.Flags().Bool("patch-text", false, "Git-apply-compatible unified patch (`git diff --binary` shape); not combinable with `--json`")
	diffCmd.Flags().Bool("quiet", false, "Print nothing; implies `--exit-code`")
	diffCmd.Flags().Bool("shortstat", false, "Only the summary line (`N files changed, …`)")
	diffCmd.Flags().BoolP("staged", "S", false, "Staged changes only (index vs HEAD)")
	diffCmd.Flags().Bool("stat", false, "Summary with change bars per file (like `git diff --stat`)")
	diffCmd.Flags().StringP("unified", "U", "", "Lines of context around each change (Git `-U`)")
	diffCmd.Flags().Bool("unstaged", false, "Unstaged changes only (worktree vs index)")
	diffCmd.Flag("find-renames").NoOptDefVal = " "
	rootCmd.AddCommand(diffCmd)

	carapace.Gen(diffCmd).FlagCompletion(carapace.ActionMap{
		"against": git.ActionRefs(git.RefOption{}.Default()),
	})

	carapace.Gen(diffCmd).PositionalCompletion(
		git.ActionRefRanges(git.RefOption{}.Default()),
	)

	carapace.Gen(diffCmd).DashAnyCompletion(
		carapace.ActionPositional(diffCmd),
	)
}
