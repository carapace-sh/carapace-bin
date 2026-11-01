package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var branch_help_infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Get info about the given branch",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branch_help_infoCmd).Standalone()

	branch_helpCmd.AddCommand(branch_help_infoCmd)
}
