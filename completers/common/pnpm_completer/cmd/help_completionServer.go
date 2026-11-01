package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_completionServerCmd = &cobra.Command{
	Use:    "completion-server",
	Short:  "Dynamic completion endpoint used by generated shell scripts",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_completionServerCmd).Standalone()

	helpCmd.AddCommand(help_completionServerCmd)
}
