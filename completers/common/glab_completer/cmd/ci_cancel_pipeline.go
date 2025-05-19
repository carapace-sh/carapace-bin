package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ci_cancel_pipelineCmd = &cobra.Command{
	Use:   "pipeline <id> [flags]",
	Short: "Cancel CI/CD pipelines.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ci_cancel_pipelineCmd).Standalone()

	ci_cancel_pipelineCmd.Flags().Bool("dry-run", false, "Simulates process, but does not cancel anything.")
	ci_cancelCmd.AddCommand(ci_cancel_pipelineCmd)

	// TODO complete pipeline ids
}
