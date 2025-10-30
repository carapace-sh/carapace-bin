package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var branch_statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Provide the current state of all applied virtual branches",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branch_statusCmd).Standalone()

	branch_statusCmd.Flags().BoolP("help", "h", false, "Print help")
	branchCmd.AddCommand(branch_statusCmd)
}
