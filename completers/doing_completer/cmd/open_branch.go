package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var open_branchCmd = &cobra.Command{
	Use:   "branch",
	Short: "Open a specific BRANCH_NAME",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(open_branchCmd).Standalone()

	open_branchCmd.Flags().Bool("help", false, "Show this message and exit.")
	openCmd.AddCommand(open_branchCmd)
}
