package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_approveBuildsCmd = &cobra.Command{
	Use:   "approve-builds",
	Short: "Approve dependencies for running scripts during installation",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_approveBuildsCmd).Standalone()

	helpCmd.AddCommand(help_approveBuildsCmd)
}
