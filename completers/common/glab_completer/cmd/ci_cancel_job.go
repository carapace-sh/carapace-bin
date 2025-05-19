package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ci_cancel_jobCmd = &cobra.Command{
	Use:   "job <id> [flags]",
	Short: "Cancel CI/CD jobs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ci_cancel_jobCmd).Standalone()

	ci_cancel_jobCmd.Flags().Bool("dry-run", false, "Simulates process, but does not cancel anything.")
	ci_cancelCmd.AddCommand(ci_cancel_jobCmd)

	// TODO complete job ids
}
