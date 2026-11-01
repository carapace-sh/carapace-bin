package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ci_configCmd = &cobra.Command{
	Use:   "config <command> [flags]",
	Short: "View and inspect GitLab CI/CD configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ci_configCmd).Standalone()

	ciCmd.AddCommand(ci_configCmd)
}
