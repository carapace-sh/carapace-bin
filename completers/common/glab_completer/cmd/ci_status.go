package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var ci_statusCmd = &cobra.Command{
	Use:     "status [flags]",
	Short:   "View a running CI/CD pipeline on current or other branch specified.",
	Aliases: []string{"stats"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ci_statusCmd).Standalone()

	ci_statusCmd.Flags().StringP("branch", "b", "", "Check pipeline status for a branch. (default current branch)")
	ci_statusCmd.Flags().BoolP("compact", "c", false, "Show status in compact format.")
	ci_statusCmd.Flags().BoolP("live", "l", false, "Show status in real time until the pipeline ends.")
	ciCmd.AddCommand(ci_statusCmd)

	carapace.Gen(ci_statusCmd).FlagCompletion(carapace.ActionMap{
		"branch": action.ActionBranches(ci_statusCmd),
	})
}
