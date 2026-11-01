package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var branch_help_latestCmd = &cobra.Command{
	Use:   "latest",
	Short: "Branch latest related commands",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branch_help_latestCmd).Standalone()

	branch_helpCmd.AddCommand(branch_help_latestCmd)
}
