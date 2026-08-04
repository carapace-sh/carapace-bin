package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var step_commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Stage and commit with LLM-generated message",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(step_commitCmd).Standalone()

	step_commitCmd.Flags().String("format", "", "Output format (text, json)")
	step_commitCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	step_commitCmd.Flags().Bool("no-verify", false, "Skip hooks")
	step_commitCmd.Flags().Bool("show-prompt", false, "Show prompt without running LLM")
	step_commitCmd.Flags().String("stage", "", "What to stage before committing [default: all]")
	step_commitCmd.Flags().BoolP("yes", "y", false, "Skip approval prompts")
	stepCmd.AddCommand(step_commitCmd)

	carapace.Gen(step_commitCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("text", "json"),
		"stage": carapace.ActionValuesDescribed(
			"all", "Stage everything: untracked files + unstaged tracked changes",
			"tracked", "Stage tracked changes only (like git add -u)",
			"none", "Stage nothing, commit only what's already in the index",
		).StyleF(style.ForKeyword),
	})
}
