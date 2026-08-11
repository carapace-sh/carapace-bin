package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var branch_help_archiveCmd = &cobra.Command{
	Use:   "archive",
	Short: "Archive an existing branch",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branch_help_archiveCmd).Standalone()

	branch_helpCmd.AddCommand(branch_help_archiveCmd)
}
