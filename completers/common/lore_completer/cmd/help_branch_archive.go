package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_branch_archiveCmd = &cobra.Command{
	Use:   "archive",
	Short: "Archive an existing branch",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_branch_archiveCmd).Standalone()

	help_branchCmd.AddCommand(help_branch_archiveCmd)
}
