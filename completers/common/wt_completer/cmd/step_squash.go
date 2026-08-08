package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/wt"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var step_squashCmd = &cobra.Command{
	Use:   "squash",
	Short: "Squash commits since branching",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(step_squashCmd).Standalone()

	step_squashCmd.Flags().Bool("dry-run", false, "Preview prompt, command, and generated message without squashing")
	step_squashCmd.Flags().String("format", "", "Output format (text, json)")
	step_squashCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	step_squashCmd.Flags().Bool("no-hooks", false, "Skip hooks")
	step_squashCmd.Flags().Bool("no-verify", false, "Skip hooks (deprecated alias for --no-hooks)")
	step_squashCmd.Flags().Bool("show-prompt", false, "Render prompt to stdout without running LLM")
	step_squashCmd.Flags().String("stage", "", "What to stage before committing [default: all]")
	step_squashCmd.Flags().BoolP("yes", "y", false, "Skip approval prompts")
	stepCmd.AddCommand(step_squashCmd)

	carapace.Gen(step_squashCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("text", "json"),
		"stage": carapace.ActionValuesDescribed(
			"all", "Stage everything: untracked files + unstaged tracked changes",
			"tracked", "Stage tracked changes only (like git add -u)",
			"none", "Stage nothing, commit only what's already in the index",
		).StyleF(style.ForKeyword),
	})

	carapace.Gen(step_squashCmd).PositionalCompletion(
		wt.ActionBranches(),
	)
}
