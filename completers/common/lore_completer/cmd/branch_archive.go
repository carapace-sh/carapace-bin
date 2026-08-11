package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var branch_archiveCmd = &cobra.Command{
	Use:   "archive",
	Short: "Archive an existing branch",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branch_archiveCmd).Standalone()

	branch_archiveCmd.Flags().BoolP("help", "h", false, "Print help")
	branchCmd.AddCommand(branch_archiveCmd)
}
