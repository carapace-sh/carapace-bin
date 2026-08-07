package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var runnerCmd = &cobra.Command{
	Use:   "runner <command> [flags]",
	Short: "Manage GitLab CI/CD runners.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runnerCmd).Standalone()

	rootCmd.AddCommand(runnerCmd)
}
