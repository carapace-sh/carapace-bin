package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var ci_statusCmd = &cobra.Command{
	Use:     "status [flags]",
	Short:   "View CI/CD pipeline status.",
	Aliases: []string{"stats"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ci_statusCmd).Standalone()

	ci_statusCmd.Flags().StringP("branch", "b", "", "Check pipeline status for a branch. Defaults to the current branch.")
	ci_statusCmd.Flags().BoolP("compact", "c", false, "Show status in compact format.")
	ci_statusCmd.Flags().String("jq", "", "Filter JSON output with a jq expression.")
	ci_statusCmd.Flags().BoolP("live", "l", false, "Show status in real time until the pipeline ends.")
	ci_statusCmd.Flags().StringP("output", "F", "", "Format output as: text, json. Note: JSON output is not compatible with --live, --wait, or --compact flags.")
	ci_statusCmd.Flags().BoolP("wait", "w", false, "Wait to return until the pipeline is finished, and provide output without a prompt.")
	ciCmd.AddCommand(ci_statusCmd)

	carapace.Gen(ci_statusCmd).FlagCompletion(carapace.ActionMap{
		"branch": action.ActionBranches(ci_statusCmd),
	})
}
