package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var completionBashCmd = &cobra.Command{
	Use:   "bash",
	Short: "Generate Bash completion",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(completionBashCmd).Standalone()

	completionCmd.AddCommand(completionBashCmd)
}
