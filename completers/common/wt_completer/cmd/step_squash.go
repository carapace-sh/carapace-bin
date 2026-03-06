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

	step_squashCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	step_squashCmd.Flags().Bool("no-verify", false, "Skip hooks")
	step_squashCmd.Flags().Bool("show-prompt", false, "Show prompt without running LLM")
	step_squashCmd.Flags().String("stage", "", "What to stage before committing [default: all]")
	step_squashCmd.Flags().BoolP("yes", "y", false, "Skip approval prompts")
	stepCmd.AddCommand(step_squashCmd)

	carapace.Gen(step_squashCmd).FlagCompletion(carapace.ActionMap{
		"stage": carapace.ActionValues("all", "tracked", "none").StyleF(style.ForKeyword),
	})

	carapace.Gen(step_squashCmd).PositionalCompletion(
		wt.ActionBranches(), // TODO verify
	)
}
